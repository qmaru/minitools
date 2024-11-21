package blake3

import (
	"encoding/base64"
	"encoding/hex"

	"lukechampine.com/blake3"
)

type Blake3Basic struct {
	data []byte
}

func (b *Blake3Basic) Sum256(s []byte) *Blake3Basic {
	b32 := blake3.Sum256(s)
	b.data = b32[:]
	return b
}

func (b *Blake3Basic) Sum512(s []byte) *Blake3Basic {
	b64 := blake3.Sum512(s)
	b.data = b64[:]
	return b
}

func (b *Blake3Basic) ToBase64() string {
	return base64.StdEncoding.EncodeToString(b.data[:])
}

func (b *Blake3Basic) ToHex() string {
	return hex.EncodeToString(b.data[:])
}

func New() *Blake3Basic {
	return new(Blake3Basic)
}
