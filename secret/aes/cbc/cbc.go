package cbc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// AESCBCBasic
type AESCBCBasic struct{}

func New() *AESCBCBasic {
	return new(AESCBCBasic)
}

func (aess *AESCBCBasic) checkKey(key []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("invalid key size, must be 16 bytes, 24 bytes, 32bytes")
	}
	return key, nil
}

// Encrypt
func (aess *AESCBCBasic) Encrypt(plainData []byte, key []byte, iv []byte) ([]byte, error) {
	key, err := aess.checkKey(key)
	if err != nil {
		return nil, err
	}

	// Group key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// key block size
	blockSize := block.BlockSize()
	// padding code
	padding := blockSize - len(plainData)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	paddingData := append(plainData, padtext...)
	// check iv
	if iv == nil {
		iv = key[:blockSize]
	} else if len(iv) != blockSize {
		return nil, errors.New("IV Error: IV length must equal block size")
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// create text array
	cypherData := make([]byte, len(paddingData))
	// encrypt
	blockMode.CryptBlocks(cypherData, paddingData)
	return cypherData, nil
}

// Decrypt
func (aess *AESCBCBasic) Decrypt(cypherData []byte, key []byte, iv []byte) ([]byte, error) {
	key, err := aess.checkKey(key)
	if err != nil {
		return nil, err
	}

	// Group key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// key block size
	blockSize := block.BlockSize()
	// check iv
	if iv == nil {
		iv = key[:blockSize]
	} else if len(iv) != blockSize {
		return nil, errors.New("IV Error: IV length must equal block size")
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// create text array
	plainData := make([]byte, len(cypherData))
	// decrypt
	blockMode.CryptBlocks(plainData, cypherData)
	// unpadding
	length := len(plainData)
	unpadding := int(plainData[length-1])
	plainData = plainData[:(length - unpadding)]
	return plainData, nil
}
