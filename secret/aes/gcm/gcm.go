package gcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// AESGCMBasic
type AESGCMBasic struct{}

func New() *AESGCMBasic {
	return new(AESGCMBasic)
}

func (aess *AESGCMBasic) checkKey(key []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("invalid key size, must be 16 bytes, 24 bytes, 32bytes")
	}
	return key, nil
}

func (aess *AESGCMBasic) Encrypt(plainData, key []byte) ([]byte, error) {
	key, err := aess.checkKey(key)
	if err != nil {
		return nil, err
	}

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
	key, err := aess.checkKey(key)
	if err != nil {
		return nil, err
	}
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
		return nil, errors.New("ciphertext too short")
	}

	nonce, cipherTextBytes := cipherData[:nonceSize], cipherData[nonceSize:]
	plainData, err := aesGCM.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return nil, err
	}

	return plainData, nil
}
