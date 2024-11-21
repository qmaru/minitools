package xor

import (
	"encoding/base64"
	"encoding/hex"
)

type XorBasic struct{}

func (x *XorBasic) Cipher(data []byte, key []byte) []byte {
	var result []byte
	keyLen := len(key)
	for i := 0; i < len(data); i++ {
		result = append(result, data[i]^key[i%keyLen])
	}
	return result
}

func (x *XorBasic) ToString(secret []byte) string {
	return string(secret)
}

func (x *XorBasic) ToByte(secret string) []byte {
	return []byte(secret)
}

func (x *XorBasic) ToBase64(secret []byte) string {
	return base64.StdEncoding.EncodeToString(secret)
}

func (x *XorBasic) FromBase64(b64 string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func (x *XorBasic) ToHex(secret []byte) string {
	return hex.EncodeToString(secret)
}

func (x *XorBasic) FromHex(hexStr string) ([]byte, error) {
	decoded, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func New() *XorBasic {
	return new(XorBasic)
}
