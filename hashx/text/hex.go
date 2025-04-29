package text

import (
	"encoding/hex"
)

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
