package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// Encrypt
func (aess *AESSuiteCBCBasic) Encrypt(plainData []byte, key []byte, iv []byte) ([]byte, error) {
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
func (aess *AESSuiteCBCBasic) Decrypt(cypherData []byte, key []byte, iv []byte) ([]byte, error) {
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
