package text

import (
	"encoding/base64"
)

func (t *TextBasic) base64Encode(enc *base64.Encoding, s []byte) string {
	return enc.EncodeToString(s)
}

func (t *TextBasic) Base64Encode(s []byte) string {
	return t.base64Encode(base64.StdEncoding, s)
}

func (t *TextBasic) URLBase64Encode(s []byte) string {
	return t.base64Encode(base64.URLEncoding, s)
}

func (t *TextBasic) RawBase64Encode(s []byte) string {
	return t.base64Encode(base64.RawStdEncoding, s)
}

func (t *TextBasic) RawURLBase64Encode(s []byte) string {
	return t.base64Encode(base64.RawURLEncoding, s)
}

func (t *TextBasic) base64Decoding(enc *base64.Encoding, s string) (*TextBasic, error) {
	ds, err := enc.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return &TextBasic{data: ds}, nil
}

func (t *TextBasic) Base64Decoding(s string) (*TextBasic, error) {
	return t.base64Decoding(base64.StdEncoding, s)
}

func (t *TextBasic) URLBase64Decoding(s string) (*TextBasic, error) {
	return t.base64Decoding(base64.URLEncoding, s)
}

func (t *TextBasic) RawBase64Decoding(s string) (*TextBasic, error) {
	return t.base64Decoding(base64.RawStdEncoding, s)
}

func (t *TextBasic) RawURLBase64Decoding(s string) (*TextBasic, error) {
	return t.base64Decoding(base64.RawURLEncoding, s)
}
