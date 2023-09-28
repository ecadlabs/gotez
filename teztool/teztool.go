package teztool

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/client"
	"github.com/ecadlabs/gotez/v2/crypt"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/latest"
)

type Signer interface {
	Sign(context.Context, []byte) (signature tz.Signature, err error)
}

type LocalSigner struct {
	crypt.PrivateKey
}

func (s LocalSigner) Sign(_ context.Context, message []byte) (signature tz.Signature, err error) {
	sig, err := s.PrivateKey.Sign(message)
	if err != nil {
		return nil, err
	}
	return sig.ToProtocol(), nil
}

func NewLocalSigner(priv crypt.PrivateKey) LocalSigner {
	return LocalSigner{PrivateKey: priv}
}

type Tool struct {
	Client      *client.Client
	ChainID     *tz.ChainID
	DebugLogger client.Logger
}

func New(client *client.Client, chain *tz.ChainID) *Tool {
	return &Tool{
		Client:  client,
		ChainID: chain,
	}
}

type fillAttrs struct {
	fillFee          bool
	fillCounter      bool
	fillGasLimit     bool
	fillStorageLimit bool
	proto            *tz.ProtocolHash
}

type FillAttr func(*fillAttrs)

func FillFee(a *fillAttrs)          { a.fillFee = true }
func FillCounter(a *fillAttrs)      { a.fillCounter = true }
func FillGasLimit(a *fillAttrs)     { a.fillGasLimit = true }
func FillStorageLimit(a *fillAttrs) { a.fillStorageLimit = true }
func FillAll(a *fillAttrs) {
	a.fillFee = true
	a.fillCounter = true
	a.fillGasLimit = true
	a.fillStorageLimit = true
}

func proto(p *tz.ProtocolHash) func(*fillAttrs) {
	return func(fa *fillAttrs) { fa.proto = p }
}

func incCounter(x tz.BigUint) tz.BigUint {
	i := x.Int()
	i.Add(i, big.NewInt(1))
	out, _ := tz.NewBigUint(i)
	return out
}

var (
	gasSafetyMargin = big.NewInt(100)
	// https://gitlab.com/tezos/tezos/-/blob/master/src/proto_alpha/lib_delegate/baking_configuration.ml#L99
	minimalFeesMutez         = big.NewInt(100)
	minimalMutezPerByte      = big.NewInt(1)
	minimalNanotezPerGasUnit = big.NewInt(100)
	storageSafetyMargin      = big.NewInt(20)
)

func mustBigUint(x *big.Int) tz.BigUint {
	v, err := tz.NewBigUint(x)
	if err != nil {
		panic(err)
	}
	return v
}

func (t *Tool) debug(format string, a ...any) {
	if t.DebugLogger != nil {
		t.DebugLogger.Printf(format, a...)
	}
}

func (t *Tool) Fill(ctx context.Context, group *latest.UnsignedOperation, attributes ...FillAttr) error {
	var attr fillAttrs
	for _, a := range attributes {
		a(&attr)
	}

	if !attr.fillFee && !attr.fillGasLimit && !attr.fillStorageLimit && !attr.fillCounter {
		return nil
	}

	var proto *tz.ProtocolHash
	if attr.proto != nil {
		proto = attr.proto
	} else {
		blockInfo, err := t.Client.BasicBlockInfo(ctx, t.ChainID.String(), group.Branch.String())
		if err != nil {
			return err
		}
		proto = blockInfo.Protocol
	}

	// fill counters
	if attr.fillCounter {
		counters := make(map[string]tz.BigUint)
		for _, op := range group.Contents {
			if op, ok := op.(core.ManagerOperation); ok {
				src := op.GetSource()
				if id, ok := src.(core.ContractID); ok {
					counter, ok := counters[id.String()]
					if !ok {
						var err error
						counter, err = t.Client.ContractCounter(ctx, &client.ContractRequest{
							Chain: t.ChainID.String(),
							Block: group.Branch.String(),
							ID:    id,
						})
						if err != nil {
							return err
						}
						t.debug("teztool: %v counter = %v", id, counter)
					}
					counter = incCounter(counter)
					counters[id.String()] = counter
					op.SetCounter(counter)
				}
			}
		}
	}

	if !attr.fillGasLimit && !attr.fillStorageLimit && !attr.fillFee {
		return nil
	}

	// get constants
	t.debug("teztool: getting constants for %v", group.Branch)
	constants, err := t.Client.Constants(ctx, &client.ContextRequest{
		Chain:    t.ChainID.String(),
		Block:    group.Branch.String(),
		Protocol: proto,
	})
	if err != nil {
		return err
	}

	// count actual manager ops
	var managerOpsCount int
	for _, op := range group.Contents {
		if _, ok := op.(core.ManagerOperation); ok {
			managerOpsCount++
		}
	}

	// set limits for dry run
	hardGasLimit := constants.GetHardGasLimitPerOperation().Int()
	hardGasLimitBlockByOp := new(big.Int).Div(constants.GetHardGasLimitPerBlock().Int(), big.NewInt(int64(managerOpsCount)))
	if hardGasLimitBlockByOp.Cmp(hardGasLimit) < 0 {
		hardGasLimit = hardGasLimitBlockByOp
	}
	for _, op := range group.Contents {
		if op, ok := op.(core.ManagerOperation); ok {
			if attr.fillGasLimit {
				op.SetGasLimit(mustBigUint(hardGasLimit))
			}
			if attr.fillStorageLimit {
				op.SetStorageLimit(mustBigUint(constants.GetHardStorageLimitPerOperation().Int()))
			}
			if attr.fillFee {
				op.SetFee(mustBigUint(minimalFeesMutez))
			}
		}
	}

	groupZeroSig := latest.SignedOperation{
		UnsignedOperation: *group,
		Signature:         &tz.GenericSignature{},
	}

	t.debug("teztool: dry run")
	runResult, err := t.Client.RunOperation(ctx, &client.RunOperationRequest{
		Chain: t.ChainID.String(),
		Block: group.Branch.String(),
		Payload: &latest.RunOperationRequest{
			Operation: groupZeroSig,
			ChainID:   t.ChainID,
		},
	})
	if err != nil {
		return err
	}

	if t.DebugLogger != nil {
		buf, _ := json.MarshalIndent(runResult, "", "    ")
		t.debug("teztool: dry run result: %s", string(buf))
	}

	resultOps := runResult.Operations()
	if len(resultOps) != len(group.Contents) {
		return errors.New("teztool: unexpected number of operations in reply")
	}

	t.debug("teztool: collecting milligas and storage")
	for i, op := range group.Contents {
		manager, ok := op.(core.ManagerOperation)
		if !ok {
			continue
		}

		resultGas, resultStorage := collectMilligasAndStorage(resultOps[i], constants)
		resultGas.Add(resultGas, big.NewInt(1000-1))
		resultGas.Div(resultGas, big.NewInt(1000))
		var isImplDest bool
		if tx, ok := op.(core.Transaction); ok {
			if _, ok := tx.GetDestination().(core.ImplicitContract); ok {
				isImplDest = true
			}
		}
		if op.OperationKind() != "reveal" && op.OperationKind() != "delegation" && op.OperationKind() != "increase_paid_storage" && !isImplDest {
			resultGas.Add(resultGas, gasSafetyMargin)
		}
		if resultStorage.Sign() != 0 {
			resultStorage.Add(resultStorage, storageSafetyMargin)
		}

		if attr.fillStorageLimit {
			manager.SetStorageLimit(mustBigUint(resultStorage))
		}

		var consumedGas *big.Int
		if attr.fillGasLimit {
			consumedGas = resultGas
			manager.SetGasLimit(mustBigUint(resultGas))
		} else {
			// use source op
			consumedGas = manager.GetGasLimit().Int()
		}

		// compute fee
		if attr.fillFee {
			gasFee := new(big.Int).Mul(consumedGas, minimalNanotezPerGasUnit)
			gasFee.Add(gasFee, big.NewInt(1000-1))
			gasFee.Div(gasFee, big.NewInt(1000)) // nanotez*gas to utez*gas

			for {
				dummyGrp := latest.SignedOperation{
					UnsignedOperation: latest.UnsignedOperation{
						Branch:   &tz.BlockHash{},
						Contents: []latest.OperationContents{op},
					},
					Signature: &tz.GenericSignature{},
				}
				var buf bytes.Buffer
				if err := encoding.Encode(&buf, &dummyGrp); err != nil {
					return err
				}
				opSize := buf.Len()
				sizeFee := new(big.Int).Mul(minimalMutezPerByte, big.NewInt(int64(opSize)))

				// https://gitlab.com/tezos/tezos/-/blob/master/src/proto_alpha/lib_client/injection.ml#L136
				x := new(big.Int).Add(minimalFeesMutez, sizeFee)
				x.Add(x, gasFee)

				done := x.Cmp(manager.GetFee().Int()) <= 0
				manager.SetFee(mustBigUint(x))
				if done {
					break
				}
			}
		}
	}
	return nil
}

func collectMilligasAndStorage(op core.OperationContents, constants core.Constants) (milligas, size *big.Int) {
	milligas = new(big.Int)
	size = new(big.Int)
	if withMeta, ok := op.(core.OperationContentsAndResult); ok {
		if manager, ok := withMeta.GetMetadata().(core.ManagerOperationMetadata); ok {
			if result, ok := manager.GetResult().(core.ManagerOperationResultAppliedOrBacktracked); ok {
				if consumedMilligas, ok := result.GetResultContents().(core.ResultWithConsumedMilligas); ok {
					milligas.Add(milligas, consumedMilligas.GetConsumedMilligas().Int())
				}
				if est, ok := result.GetResultContents().(core.StorageSizeEstimator); ok {
					size.Add(size, est.EstimateStorageSize(constants))
				}
			}
			// internal operations
			for _, result := range manager.GetInternalOperationResults() {
				if result, ok := result.GetResult().(core.ManagerOperationResultAppliedOrBacktracked); ok {
					if consumedMilligas, ok := result.GetResultContents().(core.ResultWithConsumedMilligas); ok {
						milligas.Add(milligas, consumedMilligas.GetConsumedMilligas().Int())
					}
					if est, ok := result.GetResultContents().(core.StorageSizeEstimator); ok {
						size.Add(size, est.EstimateStorageSize(constants))
					}
				}
			}
		}
	}
	return
}

func Sign(ctx context.Context, signer Signer, grp *latest.UnsignedOperation) (*latest.SignedOperation, error) {
	// forge operation
	operation := latest.SignedOperation{
		UnsignedOperation: *grp,
	}

	// hash the operation with magic byte added
	var signReq protocol.SignRequest = (*protocol.GenericOperationSignRequest)(&operation.UnsignedOperation)
	var signBytes bytes.Buffer
	if err := encoding.Encode(&signBytes, &signReq); err != nil {
		return nil, err
	}

	// sign operations
	sig, err := signer.Sign(ctx, signBytes.Bytes())
	if err != nil {
		return nil, err
	}

	switch sig := sig.(type) {
	case tz.ConventionalSignature:
		operation.Signature = sig.Generic()
	case *tz.BLSSignature:
		prefix, suffix := sig.Split()
		operation.Contents = append(operation.Contents, &latest.SignaturePrefix{SignaturePrefix: (*latest.BLSSignaturePrefix)(prefix)})
		operation.Signature = suffix
	default:
		panic("invalid signature")
	}
	return &operation, nil
}

func (t *Tool) scanBlock(ctx context.Context, hash *tz.BlockHash, op *tz.OperationHash) (bool, error) {
	basic, err := t.Client.BasicBlockInfo(ctx, t.ChainID.String(), hash.String())
	if err != nil {
		return false, err
	}
	t.debug("teztool: scanning block %v", hash)
	block, err := t.Client.Block(ctx, &client.BlockRequest{
		Chain:    t.ChainID.String(),
		Block:    hash.String(),
		Metadata: client.MetadataNever,
		Protocol: basic.Protocol,
	})
	if err != nil {
		return false, err
	}
	for _, list := range block.GetOperations() {
		for _, grp := range list {
			if *grp.GetHash() == *op {
				return true, nil
			}
		}
	}
	return false, nil
}

func (t *Tool) InjectAndWait(ctx context.Context, req *client.InjectOperationRequest) (*tz.OperationHash, error) {
	// open heads stream first
	headsCtx, headsCancel := context.WithCancel(ctx)
	defer headsCancel()

	stream, errCh, err := t.Client.Heads(headsCtx, &client.HeadsRequest{Chain: t.ChainID.String()})
	if err != nil {
		return nil, err
	}
	defer func() {
		headsCancel()
		for {
			select {
			case <-errCh:
				return
			case <-stream:
			}
		}
	}()
	opHash, err := t.Client.InjectOperation(ctx, req)
	if err != nil {
		return nil, err
	}
	t.debug("teztool: op hash: %v", opHash)

	// scan blocks
	for {
		select {
		case err := <-errCh:
			return nil, err

		case head := <-stream:
			ok, err := t.scanBlock(ctx, head.Hash, opHash)
			if err != nil {
				return nil, err
			}
			if ok {
				t.debug("teztool: found in %v", head.Hash)
				return opHash, nil
			}
		}
	}
}

func (t *Tool) FillSignAndInject(ctx context.Context, signer Signer, ops []latest.OperationContents, wait bool, attributes ...FillAttr) (*tz.OperationHash, error) {
	bi, err := t.Client.BasicBlockInfo(ctx, t.ChainID.String(), "head")
	if err != nil {
		return nil, err
	}
	group := latest.UnsignedOperation{
		Branch:   bi.Hash,
		Contents: ops,
	}
	t.debug("teztool: filling missing fields")
	if err = t.Fill(ctx, &group, append([]FillAttr{proto(bi.Protocol)}, attributes...)...); err != nil {
		return nil, err
	}
	t.debug("teztool: signing")
	signedGroup, err := Sign(ctx, signer, &group)
	if err != nil {
		return nil, err
	}
	if t.DebugLogger != nil {
		buf, _ := json.MarshalIndent(signedGroup, "", "    ")
		t.debug("%s", string(buf))
	}
	t.debug("teztool: encoding signed operation")
	var buf bytes.Buffer
	if err = encoding.Encode(&buf, signedGroup); err != nil {
		return nil, err
	}

	t.debug("teztool: injecting")
	req := client.InjectOperationRequest{
		Chain:   t.ChainID.String(),
		Payload: &client.InjectRequestPayload{Contents: buf.Bytes()},
	}
	if wait {
		return t.InjectAndWait(ctx, &req)
	} else {
		return t.Client.InjectOperation(ctx, &req)
	}
}
