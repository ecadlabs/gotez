package prefix

type Prefix struct {
	Len    int
	Prefix []byte
}

// Common prefixes
// See https://gitlab.com/tezos/tezos/blob/master/src/lib_crypto/base58.ml
var (
	// 32
	BlockHash                     = Prefix{Len: 32, Prefix: []byte{1, 52}}        // B(51)
	OperationHash                 = Prefix{Len: 32, Prefix: []byte{5, 116}}       // o(51)
	OperationListHash             = Prefix{Len: 32, Prefix: []byte{133, 233}}     // Lo(52)
	OperationListListHash         = Prefix{Len: 32, Prefix: []byte{29, 159, 109}} // LLo(53)
	ProtocolHash                  = Prefix{Len: 32, Prefix: []byte{2, 170}}       // P(51)
	ContextHash                   = Prefix{Len: 32, Prefix: []byte{79, 199}}      // Co(52)
	BlockMetadataHash             = Prefix{Len: 32, Prefix: []byte{234, 249}}     // bm(52)
	OperationMetadataHash         = Prefix{Len: 32, Prefix: []byte{5, 183}}       // r(51)
	OperationMetadataListHash     = Prefix{Len: 32, Prefix: []byte{134, 39}}      // Lr(52)
	OperationMetadataListListHash = Prefix{Len: 32, Prefix: []byte{29, 159, 182}} // LLr(53)

	// 20
	Ed25519PublicKeyHash   = Prefix{Len: 20, Prefix: []byte{6, 161, 159}}   // tz1(36)
	Secp256k1PublicKeyHash = Prefix{Len: 20, Prefix: []byte{6, 161, 161}}   // tz2(36)
	P256PublicKeyHash      = Prefix{Len: 20, Prefix: []byte{6, 161, 164}}   // tz3(36)
	ContractHash           = Prefix{Len: 20, Prefix: []byte{2, 90, 121}}    // KT1(36)
	BlindedPublicKeyHash   = Prefix{Len: 20, Prefix: []byte{1, 2, 49, 223}} // btz1(37)
	BLS12_381PublicKeyHash = Prefix{Len: 20, Prefix: []byte{6, 161, 166}}   // tz4(36)
	//lint:ignore U1000 As defined in Tezos code
	L2Address          = BLS12_381PublicKeyHash
	RollupAddress      = Prefix{Len: 20, Prefix: []byte{1, 128, 120, 31}}  // txr1(37)
	SmartRollupHash    = Prefix{Len: 20, Prefix: []byte{1, 118, 132, 217}} // scr1(37)
	ZkRollupHash       = Prefix{Len: 20, Prefix: []byte{1, 23, 224, 125}}
	SmartRollupAddress = Prefix{Len: 20, Prefix: []byte{6, 124, 117}} // sr1(36)

	// 16
	CryptoboxPublicKeyHash = Prefix{Len: 16, Prefix: []byte{153, 103}} // id(30)

	// 32
	Ed25519Seed           = Prefix{Len: 32, Prefix: []byte{13, 15, 58, 7}}     // edsk(54)
	Ed25519PublicKey      = Prefix{Len: 32, Prefix: []byte{13, 15, 37, 217}}   // edpk(54)
	Secp256k1SecretKey    = Prefix{Len: 32, Prefix: []byte{17, 162, 224, 201}} // spsk(54)
	P256SecretKey         = Prefix{Len: 32, Prefix: []byte{16, 81, 238, 189}}  // p2sk(54)
	BLS12_381SecretKey    = Prefix{Len: 32, Prefix: []byte{3, 150, 192, 40}}   // BLsk(54)
	ValueHash             = Prefix{Len: 32, Prefix: []byte{1, 106, 242}}       // vh(52)
	CycleNonce            = Prefix{Len: 32, Prefix: []byte{69, 220, 169}}      // nce(53)
	ScriptExpr            = Prefix{Len: 32, Prefix: []byte{13, 44, 64, 27}}    // expr(54)
	InboxHash             = Prefix{Len: 32, Prefix: []byte{79, 148, 196}}      // txi(53)
	InboxListHash         = InboxHash
	MessageHash           = Prefix{Len: 32, Prefix: []byte{79, 149, 30}}       // txm(53)
	CommitmentHash        = Prefix{Len: 32, Prefix: []byte{79, 148, 17}}       // txc(53)
	MessageResultHash     = Prefix{Len: 32, Prefix: []byte{18, 7, 206, 87}}    // txmr(54)
	MessageResultListHash = Prefix{Len: 32, Prefix: []byte{79, 146, 82}}       // txM(53)
	WithdrawListHash      = Prefix{Len: 32, Prefix: []byte{79, 150, 72}}       // txw(53)
	MumbaiSmartRollupHash = Prefix{Len: 32, Prefix: []byte{17, 165, 134, 138}} // src1(54)

	// 56
	Ed25519EncryptedSeed        = Prefix{Len: 56, Prefix: []byte{7, 90, 60, 179, 41}}    // edesk(88)
	Secp256k1EncryptedSecretKey = Prefix{Len: 56, Prefix: []byte{9, 237, 241, 174, 150}} // spesk(88)
	P256EncryptedSecretKey      = Prefix{Len: 56, Prefix: []byte{9, 48, 57, 115, 171}}   // p2esk(88)
	BLS12_381EncryptedSecretKey = Prefix{Len: 56, Prefix: []byte{2, 5, 30, 53, 25}}      // BLesk(88)

	// 60
	Secp256k1EncryptedScalar = Prefix{Len: 60, Prefix: []byte{1, 131, 36, 86, 248}} // seesk(93)

	// 33
	Secp256k1PublicKey = Prefix{Len: 33, Prefix: []byte{3, 254, 226, 86}}  // sppk(55)
	P256PublicKey      = Prefix{Len: 33, Prefix: []byte{3, 178, 139, 127}} // p2pk(55)
	Secp256k1Scalar    = Prefix{Len: 33, Prefix: []byte{38, 248, 136}}     // SSp(53)
	Secp256k1Element   = Prefix{Len: 33, Prefix: []byte{5, 92, 0}}         // GSp(54)

	// 64
	Ed25519SecretKey   = Prefix{Len: 64, Prefix: []byte{43, 246, 78, 7}}       // edsk(98)
	Ed25519Signature   = Prefix{Len: 64, Prefix: []byte{9, 245, 205, 134, 18}} // edsig(99)
	Secp256k1Signature = Prefix{Len: 64, Prefix: []byte{13, 115, 101, 19, 63}} // spsig1(99)
	P256Signature      = Prefix{Len: 64, Prefix: []byte{54, 240, 44, 52}}      // p2sig(98)
	GenericSignature   = Prefix{Len: 64, Prefix: []byte{4, 130, 43}}           // sig(96)

	// 4
	ChainID = Prefix{Len: 4, Prefix: []byte{87, 82, 0}}

	// 169
	SaplingSpendingKey = Prefix{Len: 169, Prefix: []byte{11, 237, 20, 92}} // sask(241)

	// 43
	SaplingAddress = Prefix{Len: 43, Prefix: []byte{18, 71, 40, 223}} // zet1(69)

	// 96
	GenericAggregateSignature = Prefix{Len: 96, Prefix: []byte{2, 75, 234, 101}}  // asig(141)
	BLS12_381Signature        = Prefix{Len: 96, Prefix: []byte{40, 171, 64, 207}} // BLsig(142)

	// 48
	BLS12_381PublicKey = Prefix{Len: 48, Prefix: []byte{6, 149, 135, 204}} // BLpk(76)

	// ?
	ScCommitmentHash = Prefix{Prefix: []byte{17, 144, 21, 100}}  // scc1(54)
	ScStateHash      = Prefix{Prefix: []byte{17, 144, 122, 202}} // scs1(54)
)

// Full list of prefixes with payload lengths
var List = []*Prefix{
	&BlockHash,
	&OperationHash,
	&OperationListHash,
	&OperationListListHash,
	&ProtocolHash,
	&ContextHash,
	&BlockMetadataHash,
	&OperationMetadataHash,
	&OperationMetadataListHash,
	&OperationMetadataListListHash,
	&Ed25519PublicKeyHash,
	&Secp256k1PublicKeyHash,
	&P256PublicKeyHash,
	&ContractHash,
	&BlindedPublicKeyHash,
	&BLS12_381PublicKeyHash,
	&RollupAddress,
	&CryptoboxPublicKeyHash,
	&Ed25519Seed,
	&Ed25519PublicKey,
	&Secp256k1SecretKey,
	&P256SecretKey,
	&ValueHash,
	&CycleNonce,
	&ScriptExpr,
	&InboxHash,
	&InboxListHash,
	&MessageHash,
	&CommitmentHash,
	&MessageResultHash,
	&MessageResultListHash,
	&WithdrawListHash,
	&Ed25519EncryptedSeed,
	&Secp256k1EncryptedSecretKey,
	&P256EncryptedSecretKey,
	&Secp256k1EncryptedScalar,
	&Secp256k1PublicKey,
	&P256PublicKey,
	&Secp256k1Scalar,
	&Secp256k1Element,
	&Ed25519SecretKey,
	&Ed25519Signature,
	&Secp256k1Signature,
	&P256Signature,
	&GenericSignature,
	&ChainID,
	&SaplingSpendingKey,
	&SaplingAddress,
	&GenericAggregateSignature,
	&BLS12_381Signature,
	&BLS12_381PublicKey,
	&BLS12_381SecretKey,
	&BLS12_381EncryptedSecretKey,
	&ScCommitmentHash,
	&ScStateHash,
	&SmartRollupHash,
	&ZkRollupHash,
	&SmartRollupAddress,
	&MumbaiSmartRollupHash,
}
