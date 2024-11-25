package gcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// AESGCMBasic
type AESGCMBasic struct{}

func New() *AESGCMBasic {
	return new(AESGCMBasic)
}

func (aess *AESGCMBasic) Encrypt(plainData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherData := aesGCM.Seal(nonce, nonce, plainData, nil)
	return cipherData, nil
}

func (aess *AESGCMBasic) Decrypt(cipherData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherData) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, cipherTextBytes := cipherData[:nonceSize], cipherData[nonceSize:]
	plainData, err := aesGCM.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return nil, err
	}

	return plainData, nil
}
