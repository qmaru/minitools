package blake3

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"

	goblake3 "lukechampine.com/blake3"
)

type Blake3Basic struct {
	data   []byte
	hasher *goblake3.Hasher
	size   int
	key    []byte
}

func (b *Blake3Basic) SetSize(size int) (*Blake3Basic, error) {
	if b.hasher != nil {
		return b, errors.New("blake3: size cannot be changed after writing data")
	}
	b.size = size
	return b, nil
}

func (b *Blake3Basic) SetKey(key []byte) (*Blake3Basic, error) {
	if b.hasher != nil {
		return b, errors.New("blake3: key cannot be changed after writing data")
	}
	if len(key) != 32 {
		return b, errors.New("blake3: key length must be 32 bytes")
	}
	b.key = key
	return b, nil
}

func (b *Blake3Basic) Write(p []byte) (n int, err error) {
	if b.hasher == nil {
		if b.size == 0 {
			b.size = 32
		}
		b.hasher = goblake3.New(b.size, b.key)
	}
	return b.hasher.Write(p)
}

func (b *Blake3Basic) WriteFrom(r io.Reader) (int64, error) {
	if b.hasher == nil {
		if b.size == 0 {
			b.size = 32
		}
		b.hasher = goblake3.New(b.size, b.key)
	}
	return io.Copy(b.hasher, r)
}

func (b *Blake3Basic) Reset() *Blake3Basic {
	if b.hasher != nil {
		b.hasher.Reset()
	}
	b.data = nil
	return b
}

func (b *Blake3Basic) Sum256(s []byte) *Blake3Basic {
	b32 := goblake3.Sum256(s)
	return &Blake3Basic{data: b32[:]}
}

func (b *Blake3Basic) Sum512(s []byte) *Blake3Basic {
	b64 := goblake3.Sum512(s)
	return &Blake3Basic{data: b64[:]}
}

func (b *Blake3Basic) SumStream() *Blake3Basic {
	if b.hasher == nil {
		return &Blake3Basic{}
	}
	sum := b.hasher.Sum(nil)
	return &Blake3Basic{data: sum}
}

func (b *Blake3Basic) ToBase64() string {
	return base64.StdEncoding.EncodeToString(b.data)
}

func (s *Blake3Basic) Bytes() []byte {
	return s.data
}

func (b *Blake3Basic) ToHex() string {
	return hex.EncodeToString(b.data)
}

func New() *Blake3Basic {
	return &Blake3Basic{size: 32}
}
