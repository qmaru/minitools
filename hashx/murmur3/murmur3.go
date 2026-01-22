package murmur3

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"

	gomurmur3 "github.com/spaolacci/murmur3"
)

type Murmur3Basic struct {
	BigEndian bool
	data      []byte
	is32      bool
	is64      bool
}

func (b *Murmur3Basic) Sum32(s []byte) *Murmur3Basic {
	b32 := gomurmur3.Sum32(s)
	buf := make([]byte, 4)

	if b.BigEndian {
		binary.BigEndian.PutUint32(buf, b32)
	} else {
		binary.LittleEndian.PutUint32(buf, b32)
	}

	return &Murmur3Basic{data: buf, is32: true}
}

func (b *Murmur3Basic) Sum64(s []byte) *Murmur3Basic {
	b64 := gomurmur3.Sum64(s)
	buf := make([]byte, 8)

	if b.BigEndian {
		binary.BigEndian.PutUint64(buf, b64)
	} else {
		binary.LittleEndian.PutUint64(buf, b64)
	}

	return &Murmur3Basic{data: buf, is64: true}
}

func (b *Murmur3Basic) Sum128(s []byte) *Murmur3Basic {
	h1, h2 := gomurmur3.Sum128(s)
	h1Byte := make([]byte, 8)
	h2Byte := make([]byte, 8)

	if b.BigEndian {
		binary.BigEndian.PutUint64(h1Byte, h1)
		binary.BigEndian.PutUint64(h2Byte, h2)
	} else {
		binary.LittleEndian.PutUint64(h1Byte, h1)
		binary.LittleEndian.PutUint64(h2Byte, h2)
	}

	combined := append(h1Byte, h2Byte...)
	return &Murmur3Basic{data: combined}
}

func (b *Murmur3Basic) ToUint32() uint32 {
	if b.is32 {
		return binary.LittleEndian.Uint32(b.data)
	}
	return 0
}

func (b *Murmur3Basic) ToUint64() uint64 {
	if b.is64 {
		return binary.LittleEndian.Uint64(b.data)
	}
	return 0
}

func (b *Murmur3Basic) ToBase64() string {
	return base64.StdEncoding.EncodeToString(b.data)
}

func (b *Murmur3Basic) ToHex() string {
	return hex.EncodeToString(b.data)
}

func (s *Murmur3Basic) Bytes() []byte {
	return s.data
}

func New() *Murmur3Basic {
	return new(Murmur3Basic)
}
