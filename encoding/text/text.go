package text

import (
	"crypto/rand"
)

type TextBasic struct {
	data []byte
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
