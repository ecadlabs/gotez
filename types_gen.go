package gotez

import (
	"fmt"
	"github.com/ecadlabs/gotez/b58/base58"
	"github.com/ecadlabs/gotez/b58/prefix"
)

// Code generated by generate.go DO NOT EDIT.

type BlockHash [HashBytesLen]byte

func (self *BlockHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BlockHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self BlockHash) String() string {
	return string(self.ToBase58())
}

func (self BlockHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.BlockHash, self[:])
}

func (self *BlockHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.BlockHash {
		return fmt.Errorf("gotez: invalid BlockHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type OperationsHash [HashBytesLen]byte

func (self *OperationsHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.OperationListListHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self OperationsHash) String() string {
	return string(self.ToBase58())
}

func (self OperationsHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.OperationListListHash, self[:])
}

func (self *OperationsHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.OperationListListHash {
		return fmt.Errorf("gotez: invalid OperationsHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type ContextHash [HashBytesLen]byte

func (self *ContextHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ContextHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self ContextHash) String() string {
	return string(self.ToBase58())
}

func (self ContextHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.ContextHash, self[:])
}

func (self *ContextHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.ContextHash {
		return fmt.Errorf("gotez: invalid ContextHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type ChainID [ChainIdBytesLen]byte

func (self *ChainID) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ChainID, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self ChainID) String() string {
	return string(self.ToBase58())
}

func (self ChainID) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.ChainID, self[:])
}

func (self *ChainID) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.ChainID {
		return fmt.Errorf("gotez: invalid ChainID encoding")
	}
	copy(self[:], payload)
	return nil
}

type BlockPayloadHash [HashBytesLen]byte

func (self *BlockPayloadHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ValueHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self BlockPayloadHash) String() string {
	return string(self.ToBase58())
}

func (self BlockPayloadHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.ValueHash, self[:])
}

func (self *BlockPayloadHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.ValueHash {
		return fmt.Errorf("gotez: invalid BlockPayloadHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type CycleNonceHash [CycleNonceBytesLen]byte

func (self *CycleNonceHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.CycleNonce, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self CycleNonceHash) String() string {
	return string(self.ToBase58())
}

func (self CycleNonceHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.CycleNonce, self[:])
}

func (self *CycleNonceHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.CycleNonce {
		return fmt.Errorf("gotez: invalid CycleNonceHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type Ed25519PublicKeyHash [AddressBytesLen]byte

func (self *Ed25519PublicKeyHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Ed25519PublicKeyHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self Ed25519PublicKeyHash) String() string {
	return string(self.ToBase58())
}

func (self Ed25519PublicKeyHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.Ed25519PublicKeyHash, self[:])
}

func (self *Ed25519PublicKeyHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.Ed25519PublicKeyHash {
		return fmt.Errorf("gotez: invalid Ed25519PublicKeyHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type Secp256k1PublicKeyHash [AddressBytesLen]byte

func (self *Secp256k1PublicKeyHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Secp256k1PublicKeyHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self Secp256k1PublicKeyHash) String() string {
	return string(self.ToBase58())
}

func (self Secp256k1PublicKeyHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.Secp256k1PublicKeyHash, self[:])
}

func (self *Secp256k1PublicKeyHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.Secp256k1PublicKeyHash {
		return fmt.Errorf("gotez: invalid Secp256k1PublicKeyHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type P256PublicKeyHash [AddressBytesLen]byte

func (self *P256PublicKeyHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.P256PublicKeyHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self P256PublicKeyHash) String() string {
	return string(self.ToBase58())
}

func (self P256PublicKeyHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.P256PublicKeyHash, self[:])
}

func (self *P256PublicKeyHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.P256PublicKeyHash {
		return fmt.Errorf("gotez: invalid P256PublicKeyHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type BLSPublicKeyHash [AddressBytesLen]byte

func (self *BLSPublicKeyHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BLS12_381PublicKeyHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self BLSPublicKeyHash) String() string {
	return string(self.ToBase58())
}

func (self BLSPublicKeyHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.BLS12_381PublicKeyHash, self[:])
}

func (self *BLSPublicKeyHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.BLS12_381PublicKeyHash {
		return fmt.Errorf("gotez: invalid BLSPublicKeyHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type ProtocolHash [HashBytesLen]byte

func (self *ProtocolHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ProtocolHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self ProtocolHash) String() string {
	return string(self.ToBase58())
}

func (self ProtocolHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.ProtocolHash, self[:])
}

func (self *ProtocolHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.ProtocolHash {
		return fmt.Errorf("gotez: invalid ProtocolHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type ContractHash [AddressBytesLen]byte

func (self *ContractHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ContractHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self ContractHash) String() string {
	return string(self.ToBase58())
}

func (self ContractHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.ContractHash, self[:])
}

func (self *ContractHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.ContractHash {
		return fmt.Errorf("gotez: invalid ContractHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type Ed25519PublicKey [Ed25519PublicKeyBytesLen]byte

func (self *Ed25519PublicKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Ed25519PublicKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self Ed25519PublicKey) String() string {
	return string(self.ToBase58())
}

func (self Ed25519PublicKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.Ed25519PublicKey, self[:])
}

func (self *Ed25519PublicKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.Ed25519PublicKey {
		return fmt.Errorf("gotez: invalid Ed25519PublicKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type Secp256k1PublicKey [Secp256k1PublicKeyBytesLen]byte

func (self *Secp256k1PublicKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Secp256k1PublicKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self Secp256k1PublicKey) String() string {
	return string(self.ToBase58())
}

func (self Secp256k1PublicKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.Secp256k1PublicKey, self[:])
}

func (self *Secp256k1PublicKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.Secp256k1PublicKey {
		return fmt.Errorf("gotez: invalid Secp256k1PublicKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type P256PublicKey [P256PublicKeyBytesLen]byte

func (self *P256PublicKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.P256PublicKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self P256PublicKey) String() string {
	return string(self.ToBase58())
}

func (self P256PublicKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.P256PublicKey, self[:])
}

func (self *P256PublicKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.P256PublicKey {
		return fmt.Errorf("gotez: invalid P256PublicKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type BLSPublicKey [BLSPublicKeyBytesLen]byte

func (self *BLSPublicKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BLS12_381PublicKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self BLSPublicKey) String() string {
	return string(self.ToBase58())
}

func (self BLSPublicKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.BLS12_381PublicKey, self[:])
}

func (self *BLSPublicKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.BLS12_381PublicKey {
		return fmt.Errorf("gotez: invalid BLSPublicKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type Ed25519PrivateKey [Ed25519SeedBytesLen]byte

func (self *Ed25519PrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Ed25519Seed, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self Ed25519PrivateKey) String() string {
	return string(self.ToBase58())
}

func (self Ed25519PrivateKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.Ed25519Seed, self[:])
}

func (self *Ed25519PrivateKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.Ed25519Seed {
		return fmt.Errorf("gotez: invalid Ed25519PrivateKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type Secp256k1PrivateKey [Secp256k1PrivateKeyBytesLen]byte

func (self *Secp256k1PrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Secp256k1SecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self Secp256k1PrivateKey) String() string {
	return string(self.ToBase58())
}

func (self Secp256k1PrivateKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.Secp256k1SecretKey, self[:])
}

func (self *Secp256k1PrivateKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.Secp256k1SecretKey {
		return fmt.Errorf("gotez: invalid Secp256k1PrivateKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type P256PrivateKey [P256PrivateKeyBytesLen]byte

func (self *P256PrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.P256SecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self P256PrivateKey) String() string {
	return string(self.ToBase58())
}

func (self P256PrivateKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.P256SecretKey, self[:])
}

func (self *P256PrivateKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.P256SecretKey {
		return fmt.Errorf("gotez: invalid P256PrivateKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type BLSPrivateKey [BLSPrivateKeyBytesLen]byte

func (self *BLSPrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BLS12_381SecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self BLSPrivateKey) String() string {
	return string(self.ToBase58())
}

func (self BLSPrivateKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.BLS12_381SecretKey, self[:])
}

func (self *BLSPrivateKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.BLS12_381SecretKey {
		return fmt.Errorf("gotez: invalid BLSPrivateKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type Ed25519EncryptedPrivateKey [Ed25519EncryptedSeedBytesLen]byte

func (self *Ed25519EncryptedPrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Ed25519EncryptedSeed, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self Ed25519EncryptedPrivateKey) String() string {
	return string(self.ToBase58())
}

func (self Ed25519EncryptedPrivateKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.Ed25519EncryptedSeed, self[:])
}

func (self *Ed25519EncryptedPrivateKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.Ed25519EncryptedSeed {
		return fmt.Errorf("gotez: invalid Ed25519EncryptedPrivateKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type Secp256k1EncryptedPrivateKey [Secp256k1EncryptedPrivateKeyBytesLen]byte

func (self *Secp256k1EncryptedPrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Secp256k1EncryptedSecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self Secp256k1EncryptedPrivateKey) String() string {
	return string(self.ToBase58())
}

func (self Secp256k1EncryptedPrivateKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.Secp256k1EncryptedSecretKey, self[:])
}

func (self *Secp256k1EncryptedPrivateKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.Secp256k1EncryptedSecretKey {
		return fmt.Errorf("gotez: invalid Secp256k1EncryptedPrivateKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type P256EncryptedPrivateKey [P256EncryptedPrivateKeyBytesLen]byte

func (self *P256EncryptedPrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.P256EncryptedSecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self P256EncryptedPrivateKey) String() string {
	return string(self.ToBase58())
}

func (self P256EncryptedPrivateKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.P256EncryptedSecretKey, self[:])
}

func (self *P256EncryptedPrivateKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.P256EncryptedSecretKey {
		return fmt.Errorf("gotez: invalid P256EncryptedPrivateKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type BLSEncryptedPrivateKey [BLSEncryptedPrivateKeyBytesLen]byte

func (self *BLSEncryptedPrivateKey) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BLS12_381EncryptedSecretKey, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self BLSEncryptedPrivateKey) String() string {
	return string(self.ToBase58())
}

func (self BLSEncryptedPrivateKey) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.BLS12_381EncryptedSecretKey, self[:])
}

func (self *BLSEncryptedPrivateKey) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.BLS12_381EncryptedSecretKey {
		return fmt.Errorf("gotez: invalid BLSEncryptedPrivateKey encoding")
	}
	copy(self[:], payload)
	return nil
}

type GenericSignature [GenericSignatureBytesLen]byte

func (self *GenericSignature) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.GenericSignature, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self GenericSignature) String() string {
	return string(self.ToBase58())
}

func (self GenericSignature) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.GenericSignature, self[:])
}

func (self *GenericSignature) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.GenericSignature {
		return fmt.Errorf("gotez: invalid GenericSignature encoding")
	}
	copy(self[:], payload)
	return nil
}

type Ed25519Signature [GenericSignatureBytesLen]byte

func (self *Ed25519Signature) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Ed25519Signature, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self Ed25519Signature) String() string {
	return string(self.ToBase58())
}

func (self Ed25519Signature) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.Ed25519Signature, self[:])
}

func (self *Ed25519Signature) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.Ed25519Signature {
		return fmt.Errorf("gotez: invalid Ed25519Signature encoding")
	}
	copy(self[:], payload)
	return nil
}

type Secp256k1Signature [GenericSignatureBytesLen]byte

func (self *Secp256k1Signature) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.Secp256k1Signature, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self Secp256k1Signature) String() string {
	return string(self.ToBase58())
}

func (self Secp256k1Signature) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.Secp256k1Signature, self[:])
}

func (self *Secp256k1Signature) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.Secp256k1Signature {
		return fmt.Errorf("gotez: invalid Secp256k1Signature encoding")
	}
	copy(self[:], payload)
	return nil
}

type P256Signature [GenericSignatureBytesLen]byte

func (self *P256Signature) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.P256Signature, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self P256Signature) String() string {
	return string(self.ToBase58())
}

func (self P256Signature) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.P256Signature, self[:])
}

func (self *P256Signature) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.P256Signature {
		return fmt.Errorf("gotez: invalid P256Signature encoding")
	}
	copy(self[:], payload)
	return nil
}

type BLSSignature [BLSSignatureBytesLen]byte

func (self *BLSSignature) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BLS12_381Signature, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self BLSSignature) String() string {
	return string(self.ToBase58())
}

func (self BLSSignature) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.BLS12_381Signature, self[:])
}

func (self *BLSSignature) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.BLS12_381Signature {
		return fmt.Errorf("gotez: invalid BLSSignature encoding")
	}
	copy(self[:], payload)
	return nil
}

type BlindedPublicKeyHash [AddressBytesLen]byte

func (self *BlindedPublicKeyHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.BlindedPublicKeyHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self BlindedPublicKeyHash) String() string {
	return string(self.ToBase58())
}

func (self BlindedPublicKeyHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.BlindedPublicKeyHash, self[:])
}

func (self *BlindedPublicKeyHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.BlindedPublicKeyHash {
		return fmt.Errorf("gotez: invalid BlindedPublicKeyHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type TXRollupAddress [AddressBytesLen]byte

func (self *TXRollupAddress) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.TXRollupAddress, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self TXRollupAddress) String() string {
	return string(self.ToBase58())
}

func (self TXRollupAddress) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.TXRollupAddress, self[:])
}

func (self *TXRollupAddress) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.TXRollupAddress {
		return fmt.Errorf("gotez: invalid TXRollupAddress encoding")
	}
	copy(self[:], payload)
	return nil
}

type ZkRollupAddress [AddressBytesLen]byte

func (self *ZkRollupAddress) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ZkRollupHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self ZkRollupAddress) String() string {
	return string(self.ToBase58())
}

func (self ZkRollupAddress) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.ZkRollupHash, self[:])
}

func (self *ZkRollupAddress) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.ZkRollupHash {
		return fmt.Errorf("gotez: invalid ZkRollupAddress encoding")
	}
	copy(self[:], payload)
	return nil
}

type DALCommitment [SlotHeaderBytesLen]byte

func (self *DALCommitment) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.SlotHeader, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self DALCommitment) String() string {
	return string(self.ToBase58())
}

func (self DALCommitment) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.SlotHeader, self[:])
}

func (self *DALCommitment) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.SlotHeader {
		return fmt.Errorf("gotez: invalid DALCommitment encoding")
	}
	copy(self[:], payload)
	return nil
}

type ScriptExprHash [HashBytesLen]byte

func (self *ScriptExprHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ScriptExpr, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self ScriptExprHash) String() string {
	return string(self.ToBase58())
}

func (self ScriptExprHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.ScriptExpr, self[:])
}

func (self *ScriptExprHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.ScriptExpr {
		return fmt.Errorf("gotez: invalid ScriptExprHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type SmartRollupAddress [AddressBytesLen]byte

func (self *SmartRollupAddress) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.SmartRollupHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self SmartRollupAddress) String() string {
	return string(self.ToBase58())
}

func (self SmartRollupAddress) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.SmartRollupHash, self[:])
}

func (self *SmartRollupAddress) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.SmartRollupHash {
		return fmt.Errorf("gotez: invalid SmartRollupAddress encoding")
	}
	copy(self[:], payload)
	return nil
}

type SmartRollupStateHash [HashBytesLen]byte

func (self *SmartRollupStateHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.SmartRollupStateHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self SmartRollupStateHash) String() string {
	return string(self.ToBase58())
}

func (self SmartRollupStateHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.SmartRollupStateHash, self[:])
}

func (self *SmartRollupStateHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.SmartRollupStateHash {
		return fmt.Errorf("gotez: invalid SmartRollupStateHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type SmartRollupCommitmentHash [HashBytesLen]byte

func (self *SmartRollupCommitmentHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.SmartRollupCommitmentHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self SmartRollupCommitmentHash) String() string {
	return string(self.ToBase58())
}

func (self SmartRollupCommitmentHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.SmartRollupCommitmentHash, self[:])
}

func (self *SmartRollupCommitmentHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.SmartRollupCommitmentHash {
		return fmt.Errorf("gotez: invalid SmartRollupCommitmentHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type ScRollupAddress [AddressBytesLen]byte

func (self *ScRollupAddress) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ScRollupHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self ScRollupAddress) String() string {
	return string(self.ToBase58())
}

func (self ScRollupAddress) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.ScRollupHash, self[:])
}

func (self *ScRollupAddress) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.ScRollupHash {
		return fmt.Errorf("gotez: invalid ScRollupAddress encoding")
	}
	copy(self[:], payload)
	return nil
}

type ScRollupStateHash [HashBytesLen]byte

func (self *ScRollupStateHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ScRollupStateHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self ScRollupStateHash) String() string {
	return string(self.ToBase58())
}

func (self ScRollupStateHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.ScRollupStateHash, self[:])
}

func (self *ScRollupStateHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.ScRollupStateHash {
		return fmt.Errorf("gotez: invalid ScRollupStateHash encoding")
	}
	copy(self[:], payload)
	return nil
}

type ScRollupCommitmentHash [HashBytesLen]byte

func (self *ScRollupCommitmentHash) ToBase58() []byte {
	out, err := base58.EncodeTZ(&prefix.ScRollupCommitmentHash, self[:])
	if err != nil {
		panic(err)
	}
	return out
}

func (self ScRollupCommitmentHash) String() string {
	return string(self.ToBase58())
}

func (self ScRollupCommitmentHash) MarshalText() ([]byte, error) {
	return base58.EncodeTZ(&prefix.ScRollupCommitmentHash, self[:])
}

func (self *ScRollupCommitmentHash) UnmarshalText(src []byte) error {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return err
	}
	if pre != &prefix.ScRollupCommitmentHash {
		return fmt.Errorf("gotez: invalid ScRollupCommitmentHash encoding")
	}
	copy(self[:], payload)
	return nil
}

