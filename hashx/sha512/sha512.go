package sha512

import (
	gosha512 "crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"hash"
	"io"
)

type SHA512Basic struct {
	data    []byte
	hasher  hash.Hash
	variant int // 384 or 512
}

func (s *SHA512Basic) SetSize(size int) (*SHA512Basic, error) {
	if s.hasher != nil {
		return s, errors.New("sha512: size cannot be changed after writing data")
	}
	if size == 0 {
		s.variant = 512
		return s, nil
	}
	if size != 384 && size != 512 {
		return s, errors.New("sha512: invalid size, supported 384 or 512")
	}
	s.variant = size
	return s, nil
}

func (s *SHA512Basic) initHasherIfNeeded() {
	if s.hasher != nil {
		return
	}
	if s.variant == 384 {
		s.hasher = gosha512.New384()
	} else {
		s.hasher = gosha512.New()
	}
}

func (s *SHA512Basic) Write(p []byte) (n int, err error) {
	s.initHasherIfNeeded()
	return s.hasher.Write(p)
}

func (s *SHA512Basic) WriteFrom(r io.Reader) (int64, error) {
	s.initHasherIfNeeded()
	return io.Copy(s.hasher, r)
}

func (s *SHA512Basic) Reset() *SHA512Basic {
	if s.hasher != nil {
		s.hasher.Reset()
	}
	s.data = nil
	return s
}

func (s *SHA512Basic) Sum512(b []byte) *SHA512Basic {
	sum := gosha512.Sum512(b)
	return &SHA512Basic{data: sum[:], variant: 512}
}

func (s *SHA512Basic) Sum384(b []byte) *SHA512Basic {
	sum := gosha512.Sum384(b)
	return &SHA512Basic{data: sum[:], variant: 384}
}

func (s *SHA512Basic) SumStream() *SHA512Basic {
	if s.hasher == nil {
		return &SHA512Basic{}
	}
	sum := s.hasher.Sum(nil)
	return &SHA512Basic{data: sum, variant: s.variant}
}

func (s *SHA512Basic) ToBase64() string {
	return base64.StdEncoding.EncodeToString(s.data)
}

func (s *SHA512Basic) Bytes() []byte {
	return s.data
}

func (s *SHA512Basic) ToHex() string {
	return hex.EncodeToString(s.data)
}

func New() *SHA512Basic {
	return &SHA512Basic{variant: 512}
}
