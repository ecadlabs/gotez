package gotez

import (
	"github.com/ecadlabs/gotez/b58/base58"
	"github.com/ecadlabs/gotez/b58/prefix"
)

// Code generated by generate.go DO NOT EDIT.

type BlockHash [BlockHashBytesLen]byte

func (self *BlockHash) String() string {
	return string(self.ToBase58())
}

func (self *BlockHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BlockHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type OperationsHash [OperationListListHashBytesLen]byte

func (self *OperationsHash) String() string {
	return string(self.ToBase58())
}

func (self *OperationsHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.OperationListListHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type ContextHash [ContextHashBytesLen]byte

func (self *ContextHash) String() string {
	return string(self.ToBase58())
}

func (self *ContextHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ContextHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type ChainID [ChainIdBytesLen]byte

func (self *ChainID) String() string {
	return string(self.ToBase58())
}

func (self *ChainID) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ChainID, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type BlockPayloadHash [BlockPayloadHashBytesLen]byte

func (self *BlockPayloadHash) String() string {
	return string(self.ToBase58())
}

func (self *BlockPayloadHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ValueHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type CycleNonceHash [CycleNonceBytesLen]byte

func (self *CycleNonceHash) String() string {
	return string(self.ToBase58())
}

func (self *CycleNonceHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.CycleNonce, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type Ed25519PublicKeyHash [PKHBytesLen]byte

func (self *Ed25519PublicKeyHash) String() string {
	return string(self.ToBase58())
}

func (self *Ed25519PublicKeyHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Ed25519PublicKeyHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type Secp256k1PublicKeyHash [PKHBytesLen]byte

func (self *Secp256k1PublicKeyHash) String() string {
	return string(self.ToBase58())
}

func (self *Secp256k1PublicKeyHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Secp256k1PublicKeyHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type P256PublicKeyHash [PKHBytesLen]byte

func (self *P256PublicKeyHash) String() string {
	return string(self.ToBase58())
}

func (self *P256PublicKeyHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.P256PublicKeyHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type BLSPublicKeyHash [PKHBytesLen]byte

func (self *BLSPublicKeyHash) String() string {
	return string(self.ToBase58())
}

func (self *BLSPublicKeyHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BLS12_381PublicKeyHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type ProtocolHash [ProtocolHashBytesLen]byte

func (self *ProtocolHash) String() string {
	return string(self.ToBase58())
}

func (self *ProtocolHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ProtocolHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type ContractHash [ContractHashBytesLen]byte

func (self *ContractHash) String() string {
	return string(self.ToBase58())
}

func (self *ContractHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ContractHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type Ed25519PublicKey [Ed25519PublicKeyBytesLen]byte

func (self *Ed25519PublicKey) String() string {
	return string(self.ToBase58())
}

func (self *Ed25519PublicKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Ed25519PublicKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type Secp256k1PublicKey [Secp256k1PublicKeyBytesLen]byte

func (self *Secp256k1PublicKey) String() string {
	return string(self.ToBase58())
}

func (self *Secp256k1PublicKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Secp256k1PublicKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type P256PublicKey [P256PublicKeyBytesLen]byte

func (self *P256PublicKey) String() string {
	return string(self.ToBase58())
}

func (self *P256PublicKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.P256PublicKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type BLSPublicKey [BLSPublicKeyBytesLen]byte

func (self *BLSPublicKey) String() string {
	return string(self.ToBase58())
}

func (self *BLSPublicKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BLS12_381PublicKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type Ed25519PrivateKey [Ed25519SeedBytesLen]byte

func (self *Ed25519PrivateKey) String() string {
	return string(self.ToBase58())
}

func (self *Ed25519PrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Ed25519Seed, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type Secp256k1PrivateKey [Secp256k1PrivateKeyBytesLen]byte

func (self *Secp256k1PrivateKey) String() string {
	return string(self.ToBase58())
}

func (self *Secp256k1PrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Secp256k1SecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type P256PrivateKey [P256PrivateKeyBytesLen]byte

func (self *P256PrivateKey) String() string {
	return string(self.ToBase58())
}

func (self *P256PrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.P256SecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type BLSPrivateKey [BLSPrivateKeyBytesLen]byte

func (self *BLSPrivateKey) String() string {
	return string(self.ToBase58())
}

func (self *BLSPrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BLS12_381SecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type Ed25519EncryptedPrivateKey [Ed25519EncryptedSeedBytesLen]byte

func (self *Ed25519EncryptedPrivateKey) String() string {
	return string(self.ToBase58())
}

func (self *Ed25519EncryptedPrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Ed25519EncryptedSeed, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type Secp256k1EncryptedPrivateKey [Secp256k1EncryptedPrivateKeyBytesLen]byte

func (self *Secp256k1EncryptedPrivateKey) String() string {
	return string(self.ToBase58())
}

func (self *Secp256k1EncryptedPrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Secp256k1EncryptedSecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type P256EncryptedPrivateKey [P256EncryptedPrivateKeyBytesLen]byte

func (self *P256EncryptedPrivateKey) String() string {
	return string(self.ToBase58())
}

func (self *P256EncryptedPrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.P256EncryptedSecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type BLSEncryptedPrivateKey [BLSEncryptedPrivateKeyBytesLen]byte

func (self *BLSEncryptedPrivateKey) String() string {
	return string(self.ToBase58())
}

func (self *BLSEncryptedPrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BLS12_381EncryptedSecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type GenericSignature [GenericSignatureBytesLen]byte

func (self *GenericSignature) String() string {
	return string(self.ToBase58())
}

func (self *GenericSignature) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.GenericSignature, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type Ed25519Signature [GenericSignatureBytesLen]byte

func (self *Ed25519Signature) String() string {
	return string(self.ToBase58())
}

func (self *Ed25519Signature) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Ed25519Signature, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type Secp256k1Signature [GenericSignatureBytesLen]byte

func (self *Secp256k1Signature) String() string {
	return string(self.ToBase58())
}

func (self *Secp256k1Signature) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Secp256k1Signature, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type P256Signature [GenericSignatureBytesLen]byte

func (self *P256Signature) String() string {
	return string(self.ToBase58())
}

func (self *P256Signature) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.P256Signature, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

type BLSSignature [BLSSignatureBytesLen]byte

func (self *BLSSignature) String() string {
	return string(self.ToBase58())
}

func (self *BLSSignature) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BLS12_381Signature, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

