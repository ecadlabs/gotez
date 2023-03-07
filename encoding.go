package gotez

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ecadlabs/gotez/base58"
)

type Base58Prefix struct {
	plLen  int
	prefix []byte
}

// Common prefixes
// See https://gitlab.com/tezos/tezos/blob/master/src/lib_crypto/base58.ml
var (
	// 32
	PfxBlockHash                     = Base58Prefix{plLen: 32, prefix: []byte{1, 52}}        // B(51)
	PfxOperationHash                 = Base58Prefix{plLen: 32, prefix: []byte{5, 116}}       // o(51)
	PfxOperationListHash             = Base58Prefix{plLen: 32, prefix: []byte{133, 233}}     // Lo(52)
	PfxOperationListListHash         = Base58Prefix{plLen: 32, prefix: []byte{29, 159, 109}} // LLo(53)
	PfxProtocolHash                  = Base58Prefix{plLen: 32, prefix: []byte{2, 170}}       // P(51)
	PfxContextHash                   = Base58Prefix{plLen: 32, prefix: []byte{79, 199}}      // Co(52)
	PfxBlockMetadataHash             = Base58Prefix{plLen: 32, prefix: []byte{234, 249}}     // bm(52)
	PfxOperationMetadataHash         = Base58Prefix{plLen: 32, prefix: []byte{5, 183}}       // r(51)
	PfxOperationMetadataListHash     = Base58Prefix{plLen: 32, prefix: []byte{134, 39}}      // Lr(52)
	PfxOperationMetadataListListHash = Base58Prefix{plLen: 32, prefix: []byte{29, 159, 182}} // LLr(53)

	// 20
	PfxEd25519PublicKeyHash   = Base58Prefix{plLen: 20, prefix: []byte{6, 161, 159}}   // tz1(36)
	PfxSecp256k1PublicKeyHash = Base58Prefix{plLen: 20, prefix: []byte{6, 161, 161}}   // tz2(36)
	PfxP256PublicKeyHash      = Base58Prefix{plLen: 20, prefix: []byte{6, 161, 164}}   // tz3(36)
	PfxContractHash           = Base58Prefix{plLen: 20, prefix: []byte{2, 90, 121}}    // KT1(36)
	PfxBlindedPublicKeyHash   = Base58Prefix{plLen: 20, prefix: []byte{1, 2, 49, 223}} // btz1(37)
	PfxBLS12_381PublicKeyHash = Base58Prefix{plLen: 20, prefix: []byte{6, 161, 166}}   // tz4(36)
	//lint:ignore U1000 As defined in Tezos code
	PfxL2Address     = PfxBLS12_381PublicKeyHash
	PfxRollupAddress = Base58Prefix{plLen: 20, prefix: []byte{1, 128, 120, 31}}  // txr1(37)
	PfxScRollupHash  = Base58Prefix{plLen: 20, prefix: []byte{1, 118, 132, 217}} // scr1(37)

	// 16
	PfxCryptoboxPublicKeyHash = Base58Prefix{plLen: 16, prefix: []byte{153, 103}} // id(30)

	// 32
	PfxEd25519Seed           = Base58Prefix{plLen: 32, prefix: []byte{13, 15, 58, 7}}     // edsk(54)
	PfxEd25519PublicKey      = Base58Prefix{plLen: 32, prefix: []byte{13, 15, 37, 217}}   // edpk(54)
	PfxSecp256k1SecretKey    = Base58Prefix{plLen: 32, prefix: []byte{17, 162, 224, 201}} // spsk(54)
	PfxP256SecretKey         = Base58Prefix{plLen: 32, prefix: []byte{16, 81, 238, 189}}  // p2sk(54)
	PfxBLS12_381SecretKey    = Base58Prefix{plLen: 32, prefix: []byte{3, 150, 192, 40}}   // BLsk(54)
	PfxValueHash             = Base58Prefix{plLen: 32, prefix: []byte{1, 106, 242}}       // vh(52)
	PfxCycleNonce            = Base58Prefix{plLen: 32, prefix: []byte{69, 220, 169}}      // nce(53)
	PfxScriptExpr            = Base58Prefix{plLen: 32, prefix: []byte{13, 44, 64, 27}}    // expr(54)
	PfxInboxHash             = Base58Prefix{plLen: 32, prefix: []byte{79, 148, 196}}      // txi(53)
	PfxInboxListHash         = PfxInboxHash
	PfxMessageHash           = Base58Prefix{plLen: 32, prefix: []byte{79, 149, 30}}    // txm(53)
	PfxCommitmentHash        = Base58Prefix{plLen: 32, prefix: []byte{79, 148, 17}}    // txc(53)
	PfxMessageResultHash     = Base58Prefix{plLen: 32, prefix: []byte{18, 7, 206, 87}} // txmr(54)
	PfxMessageResultListHash = Base58Prefix{plLen: 32, prefix: []byte{79, 146, 82}}    // txM(53)
	PfxWithdrawListHash      = Base58Prefix{plLen: 32, prefix: []byte{79, 150, 72}}    // txw(53)

	// 56
	PfxEd25519EncryptedSeed        = Base58Prefix{plLen: 56, prefix: []byte{7, 90, 60, 179, 41}}    // edesk(88)
	PfxSecp256k1EncryptedSecretKey = Base58Prefix{plLen: 56, prefix: []byte{9, 237, 241, 174, 150}} // spesk(88)
	PfxP256EncryptedSecretKey      = Base58Prefix{plLen: 56, prefix: []byte{9, 48, 57, 115, 171}}   // p2esk(88)
	PfxBLS12_381EncryptedSecretKey = Base58Prefix{plLen: 56, prefix: []byte{2, 5, 30, 53, 25}}      // BLesk(88)

	// 60
	PfxSecp256k1EncryptedScalar = Base58Prefix{plLen: 60, prefix: []byte{1, 131, 36, 86, 248}} // seesk(93)

	// 33
	PfxSecp256k1PublicKey = Base58Prefix{plLen: 33, prefix: []byte{3, 254, 226, 86}}  // sppk(55)
	PfxP256PublicKey      = Base58Prefix{plLen: 33, prefix: []byte{3, 178, 139, 127}} // p2pk(55)
	PfxSecp256k1Scalar    = Base58Prefix{plLen: 33, prefix: []byte{38, 248, 136}}     // SSp(53)
	PfxSecp256k1Element   = Base58Prefix{plLen: 33, prefix: []byte{5, 92, 0}}         // GSp(54)

	// 64
	PfxEd25519SecretKey   = Base58Prefix{plLen: 64, prefix: []byte{43, 246, 78, 7}}       // edsk(98)
	PfxEd25519Signature   = Base58Prefix{plLen: 64, prefix: []byte{9, 245, 205, 134, 18}} // edsig(99)
	PfxSecp256k1Signature = Base58Prefix{plLen: 64, prefix: []byte{13, 115, 101, 19, 63}} // spsig1(99)
	PfxP256Signature      = Base58Prefix{plLen: 64, prefix: []byte{54, 240, 44, 52}}      // p2sig(98)
	PfxGenericSignature   = Base58Prefix{plLen: 64, prefix: []byte{4, 130, 43}}           // sig(96)

	// 4
	PfxChainID = Base58Prefix{plLen: 4, prefix: []byte{87, 82, 0}}

	// 169
	PfxSaplingSpendingKey = Base58Prefix{plLen: 169, prefix: []byte{11, 237, 20, 92}} // sask(241)

	// 43
	PfxSaplingAddress = Base58Prefix{plLen: 43, prefix: []byte{18, 71, 40, 223}} // zet1(69)

	// 96
	PfxGenericAggregateSignature = Base58Prefix{plLen: 96, prefix: []byte{2, 75, 234, 101}}  // asig(141)
	PfxBLS12_381Signature        = Base58Prefix{plLen: 96, prefix: []byte{40, 171, 64, 207}} // BLsig(142)

	// 48
	PfxBLS12_381PublicKey = Base58Prefix{plLen: 48, prefix: []byte{6, 149, 135, 204}} // BLpk(76)

	// ?
	PfxScCommitmentHash = Base58Prefix{prefix: []byte{17, 144, 21, 100}}  // scc1(54)
	PfxScStateHash      = Base58Prefix{prefix: []byte{17, 144, 122, 202}} // scs1(54)
)

// Full list of prefixes with payload lengths
var commonPrefixes = []*Base58Prefix{
	&PfxBlockHash,
	&PfxOperationHash,
	&PfxOperationListHash,
	&PfxOperationListListHash,
	&PfxProtocolHash,
	&PfxContextHash,
	&PfxBlockMetadataHash,
	&PfxOperationMetadataHash,
	&PfxOperationMetadataListHash,
	&PfxOperationMetadataListListHash,
	&PfxEd25519PublicKeyHash,
	&PfxSecp256k1PublicKeyHash,
	&PfxP256PublicKeyHash,
	&PfxContractHash,
	&PfxBlindedPublicKeyHash,
	&PfxBLS12_381PublicKeyHash,
	&PfxRollupAddress,
	&PfxCryptoboxPublicKeyHash,
	&PfxEd25519Seed,
	&PfxEd25519PublicKey,
	&PfxSecp256k1SecretKey,
	&PfxP256SecretKey,
	&PfxValueHash,
	&PfxCycleNonce,
	&PfxScriptExpr,
	&PfxInboxHash,
	&PfxInboxListHash,
	&PfxMessageHash,
	&PfxCommitmentHash,
	&PfxMessageResultHash,
	&PfxMessageResultListHash,
	&PfxWithdrawListHash,
	&PfxEd25519EncryptedSeed,
	&PfxSecp256k1EncryptedSecretKey,
	&PfxP256EncryptedSecretKey,
	&PfxSecp256k1EncryptedScalar,
	&PfxSecp256k1PublicKey,
	&PfxP256PublicKey,
	&PfxSecp256k1Scalar,
	&PfxSecp256k1Element,
	&PfxEd25519SecretKey,
	&PfxEd25519Signature,
	&PfxSecp256k1Signature,
	&PfxP256Signature,
	&PfxGenericSignature,
	&PfxChainID,
	&PfxSaplingSpendingKey,
	&PfxSaplingAddress,
	&PfxGenericAggregateSignature,
	&PfxBLS12_381Signature,
	&PfxBLS12_381PublicKey,
	&PfxBLS12_381SecretKey,
	&PfxBLS12_381EncryptedSecretKey,
	&PfxScCommitmentHash,
	&PfxScStateHash,
	&PfxScRollupHash,
}

// ErrPrefix is returned in case of unknown Tezos base58 prefix
var ErrPrefix = errors.New("gotez: unknown Tezos base58 prefix")

func DecodeTZBase58(data []byte) (prefix *Base58Prefix, payload []byte, err error) {
	buf, err := base58.DecodeCheck(data)
	if err != nil {
		return
	}
	for _, p := range commonPrefixes {
		if bytes.HasPrefix(buf, p.prefix) {
			plLen := len(buf) - len(p.prefix)
			if p.plLen != 0 && plLen != p.plLen {
				return p, nil, fmt.Errorf("gotez: invalid base58 message length: expected %d, got %d", p.plLen, plLen)
			}
			return p, buf[len(p.prefix):], nil
		}
	}
	err = ErrPrefix
	return
}

func EncodeTZBase58(prefix *Base58Prefix, payload []byte) ([]byte, error) {
	if prefix.plLen != 0 && len(payload) != prefix.plLen {
		return nil, fmt.Errorf("gotez: invalid base58 message length: expected %d, got %d", prefix.plLen, len(payload))
	}
	data := make([]byte, len(prefix.prefix)+len(payload))
	copy(data, prefix.prefix)
	copy(data[len(prefix.prefix):], payload)
	return base58.EncodeCheck(data), nil
}
