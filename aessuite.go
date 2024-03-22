package minitools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// AESSuiteBasic
type AESSuiteBasic struct{}

// Encrypt
func (aess *AESSuiteBasic) Encrypt(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	// Group key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// key block size
	blockSize := block.BlockSize()
	// padding code
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	paddingData := append(plaintext, padtext...)
	// check iv
	if iv == nil {
		iv = key[:blockSize]
	} else if len(iv) != blockSize {
		return nil, errors.New("IV Error: IV length must equal block size")
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// create text array
	cyphertext := make([]byte, len(paddingData))
	// encrypt
	blockMode.CryptBlocks(cyphertext, paddingData)
	return cyphertext, nil
}

// Decrypt
func (aess *AESSuiteBasic) Decrypt(cryted []byte, key []byte, iv []byte) ([]byte, error) {
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
	plaintext := make([]byte, len(cryted))
	// decrypt
	blockMode.CryptBlocks(plaintext, cryted)
	// unpadding
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	plaintext = plaintext[:(length - unpadding)]
	return plaintext, nil
}
