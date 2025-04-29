package text

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

type TextBasic struct {
	data []byte
}

func (t *TextBasic) Base64Encode(s []byte) string {
	return base64.StdEncoding.EncodeToString(s)
}

func (t *TextBasic) Base64Decoding(s string) (*TextBasic, error) {
	ds, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return &TextBasic{data: ds}, nil
}

func (t *TextBasic) HexEncode(s []byte) string {
	return hex.EncodeToString(s)
}

func (t *TextBasic) HexDecoding(s string) (*TextBasic, error) {
	ds, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	t.data = ds
	return &TextBasic{data: ds}, nil
}

func (t *TextBasic) Nonce(l int) ([]byte, error) {
	b := make([]byte, l)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (t *TextBasic) DecodeRaw() []byte {
	return t.data
}

func (t *TextBasic) DecodeString() string {
	return string(t.data)
}

func New() *TextBasic {
	return new(TextBasic)
}
