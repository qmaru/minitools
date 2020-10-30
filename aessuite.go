package minitools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AESSuiteBasic AES 套件基类
type AESSuiteBasic struct{}

// Encrypt 加密
func (aess *AESSuiteBasic) Encrypt(plaintext []byte, key []byte, iv []byte) []byte {
	// 分组密钥
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 填充码
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	paddingData := append(plaintext, padtext...)
	// 加密模式
	if iv == nil {
		iv = key[:blockSize]
	} else if len(iv) != blockSize {
		panic("IV Error: IV length must equal block size")
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 创建加密数组
	cyphertext := make([]byte, len(paddingData))
	// 执行加密
	blockMode.CryptBlocks(cyphertext, paddingData)
	return cyphertext
}

// Decrypt 解密函数
func (aess *AESSuiteBasic) Decrypt(cryted []byte, key []byte, iv []byte) []byte {
	// 分组秘钥
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	if iv == nil {
		iv = key[:blockSize]
	} else if len(iv) != blockSize {
		panic("IV Error: IV length must equal block size")
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 创建解密数组
	plaintext := make([]byte, len(cryted))
	// 解密
	blockMode.CryptBlocks(plaintext, cryted)
	// 去码
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	plaintext = plaintext[:(length - unpadding)]
	return plaintext
}
