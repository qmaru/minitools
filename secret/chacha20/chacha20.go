package chacha20

import (
	"crypto/rand"
	"errors"

	gochacha "golang.org/x/crypto/chacha20"
)

type Chacha20Basic struct{}

func New() *Chacha20Basic {
	return new(Chacha20Basic)
}

func (c *Chacha20Basic) GenerateNonce() ([]byte, error) {
	nonce := make([]byte, 12)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

func (c *Chacha20Basic) Encrypt(plainData, key, nonce []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key size: must 32 bytes")
	}

	cipher, err := gochacha.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plainData))
	cipher.XORKeyStream(ciphertext, plainData)
	return append(nonce, ciphertext...), nil
}

func (c *Chacha20Basic) Decrypt(cipherData, key []byte) ([]byte, error) {
	if len(cipherData) < 12 {
		return nil, errors.New("ciphertext too short")
	}

	nonce := cipherData[:12]
	cipherText := cipherData[12:]

	cipher, err := gochacha.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	plainData := make([]byte, len(cipherText))
	cipher.XORKeyStream(plainData, cipherText)

	return plainData, nil
}
