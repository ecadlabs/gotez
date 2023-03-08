package gotez

import (
	"crypto/sha512"
	"errors"

	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/pbkdf2"
)

var (
	ErrPrivateKeyDecrypt = errors.New("tezos: unable to decrypt the private key")
)

const (
	encIterations = 32768
	encKeyLen     = 32
)

func decryptPrivateKey(data []byte, passCb func() ([]byte, error)) ([]byte, error) {
	if passCb == nil {
		return nil, ErrPrivateKeyDecrypt
	}
	passphrase, err := passCb()
	if err != nil {
		return nil, err
	}

	salt, box := data[:8], data[8:]
	secretboxKey := pbkdf2.Key(passphrase, salt, encIterations, encKeyLen, sha512.New)

	var (
		tmp   [32]byte
		nonce [24]byte
	)
	copy(tmp[:], secretboxKey)
	opened, ok := secretbox.Open(nil, box, &nonce, &tmp)
	if !ok {
		return nil, ErrPrivateKeyDecrypt
	}

	return opened, nil
}
