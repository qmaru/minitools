package text

import (
	"encoding/base64"
	"encoding/hex"
)

type TextBasic struct{}

func (t *TextBasic) Base64Encode(s []byte) string {
	return base64.StdEncoding.EncodeToString(s)
}

func (t *TextBasic) Base64Decode(s string) (string, error) {
	ds, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(ds), nil
}

func (t *TextBasic) HexEncode(s []byte) string {
	return hex.EncodeToString(s)
}

func (t *TextBasic) HexDecode(s string) (string, error) {
	ds, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(ds), nil
}

func New() *TextBasic {
	return new(TextBasic)
}
