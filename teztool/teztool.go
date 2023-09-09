package teztool

import (
	"bytes"
	"context"
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

type TezTool struct {
	Client  *client.Client
	Signer  Signer
	ChainID *tz.ChainID
}

type fillAttrs struct {
	fillFee          bool
	fillCounter      bool
	fillGasLimit     bool
	fillStorageLimit bool
	blockInfo        *client.BasicBlockInfo
}

type FillAttr func(*fillAttrs)

func FillFee(a *fillAttrs)          { a.fillFee = true }
func FillCounter(a *fillAttrs)      { a.fillCounter = true }
func FillGasLimit(a *fillAttrs)     { a.fillGasLimit = true }
func FillStorageLimit(a *fillAttrs) { a.fillStorageLimit = true }

func blockInfo(i *client.BasicBlockInfo) func(*fillAttrs) {
	return func(fa *fillAttrs) { fa.blockInfo = i }
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

func (t *TezTool) Fill(ctx context.Context, ops []latest.OperationContents, attributes ...FillAttr) error {
	var attr fillAttrs
	for _, a := range attributes {
		a(&attr)
	}

	if !attr.fillFee && !attr.fillGasLimit && !attr.fillStorageLimit && !attr.fillCounter {
		return nil
	}

	var blockInfo *client.BasicBlockInfo
	if attr.blockInfo != nil {
		blockInfo = attr.blockInfo
	} else {
		var err error
		if blockInfo, err = t.Client.BasicBlockInfo(ctx, t.ChainID.String(), "head"); err != nil {
			return err
		}
	}

	// fill counters
	if attr.fillCounter {
		counters := make(map[string]tz.BigUint)
		for _, op := range ops {
			if op, ok := op.(core.ManagerOperation); ok {
				src := op.GetSource()
				if id, ok := src.(core.ContractID); ok {
					counter, ok := counters[id.String()]
					if !ok {
						var err error
						counter, err = t.Client.ContractCounter(ctx, &client.ContractRequest{
							Chain: t.ChainID.String(),
							Block: blockInfo.Hash.String(),
							ID:    id,
						})
						if err != nil {
							return err
						}
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
	constants, err := t.Client.Constants(context.Background(), &client.ContextRequest{
		Chain:    t.ChainID.String(),
		Block:    blockInfo.Hash.String(),
		Protocol: blockInfo.Protocol,
	})
	if err != nil {
		return err
	}

	for _, op := range ops {
		if op, ok := op.(core.ManagerOperation); ok {
			if attr.fillGasLimit {
				op.SetGasLimit(mustBigUint(constants.GetHardGasLimitPerOperation().Int()))
			}
			if attr.fillStorageLimit {
				op.SetStorageLimit(mustBigUint(constants.GetHardStorageLimitPerOperation().Int()))
			}
			if attr.fillFee {
				op.SetFee(tz.BigUZero())
			}
		}
	}

	group := latest.SignedOperation{
		UnsignedOperation: latest.UnsignedOperation{
			Branch:   blockInfo.Hash,
			Contents: ops,
		},
		Signature: &tz.GenericSignature{},
	}

	runResult, err := t.Client.RunOperation(context.Background(), &client.RunOperationRequest{
		Chain: t.ChainID.String(),
		Block: blockInfo.Hash.String(),
		Payload: &latest.RunOperationRequest{
			Operation: group,
			ChainID:   t.ChainID,
		},
	})
	if err != nil {
		return err
	}

	resultOps := runResult.Operations()
	if len(resultOps) != len(ops) {
		return errors.New("teztool: unexpected number of operations in reply")
	}

	for i, op := range ops {
		op, ok := op.(core.ManagerOperation)
		if !ok {
			continue
		}

		resultGas, resultStorage := collectMilligasAndStorage(resultOps[i], constants)
		resultGas.Add(resultGas, big.NewInt(1000-1))
		resultGas.Div(resultGas, big.NewInt(1000))
		resultGas.Add(resultGas, gasSafetyMargin)
		if resultStorage.Sign() != 0 {
			resultStorage.Add(resultStorage, storageSafetyMargin)
		}

		if attr.fillStorageLimit {
			op.SetStorageLimit(mustBigUint(resultStorage))
		}

		var consumedGas *big.Int
		if attr.fillGasLimit {
			consumedGas = resultGas
			op.SetGasLimit(mustBigUint(resultGas))
		} else {
			// use source op
			consumedGas = op.GetGasLimit().Int()
		}

		// compute fee
		if attr.fillFee {
			gasFee := new(big.Int).Set(consumedGas)
			gasFee.Mul(gasFee, minimalNanotezPerGasUnit)
			gasFee.Add(gasFee, big.NewInt(1000-1))
			gasFee.Div(gasFee, big.NewInt(1000)) // nanotez*gas to utez*gas

			var opSize int
			for {
				var encodedOp bytes.Buffer
				if err := encoding.Encode(&encodedOp, &op); err != nil {
					return err
				}
				opSize = encodedOp.Len()
				sizeFee := new(big.Int).Set(minimalMutezPerByte)
				sizeFee.Mul(sizeFee, big.NewInt(int64(opSize)))

				// https://gitlab.com/tezos/tezos/-/blob/master/src/proto_alpha/lib_client/injection.ml#L136
				x := new(big.Int).Set(minimalFeesMutez)
				x.Add(x, sizeFee)
				x.Add(x, gasFee)

				done := x.Cmp(op.GetFee().Int()) <= 0
				op.SetFee(mustBigUint(x))
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

func (t *TezTool) SignAndInject(ctx context.Context, ops []latest.OperationContents, attributes ...FillAttr) (*tz.OperationHash, error) {
	bi, err := t.Client.BasicBlockInfo(ctx, t.ChainID.String(), "head")
	if err != nil {
		return nil, err
	}
	if err = t.Fill(ctx, ops, append([]FillAttr{blockInfo(bi)}, attributes...)...); err != nil {
		return nil, err
	}

	// forge operation
	operation := latest.SignedOperation{
		UnsignedOperation: latest.UnsignedOperation{
			Branch:   bi.Hash,
			Contents: ops,
		},
	}

	// hash the operation with magic byte added
	var signReq protocol.SignRequest = (*protocol.GenericOperationSignRequest)(&operation.UnsignedOperation)
	var signBytes bytes.Buffer
	if err := encoding.Encode(&signBytes, &signReq); err != nil {
		return nil, err
	}

	// sign operations
	sig, err := t.Signer.Sign(ctx, signBytes.Bytes())
	if err != nil {
		return nil, err
	}
	if csig, ok := sig.(tz.ConventionalSignature); ok {
		operation.Signature = csig.Generic()
	} else {
		return nil, errors.New("teztool: BLS signatures aren't supported")
	}

	var buf bytes.Buffer
	if err = encoding.Encode(&buf, &operation); err != nil {
		return nil, err
	}

	// inject
	return t.Client.InjectOperation(ctx, &client.InjectOperationRequest{
		Chain:   t.ChainID.String(),
		Payload: &client.InjectRequestPayload{Contents: buf.Bytes()},
	})
}
