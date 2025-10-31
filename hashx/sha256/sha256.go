package sha256

import (
	gosha256 "crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"hash"
	"io"
)

type SHA2Basic struct {
	data    []byte
	hasher  hash.Hash
	variant int
}

func (s *SHA2Basic) SetSize(size int) (*SHA2Basic, error) {
	if s.hasher != nil {
		return s, errors.New("sha2: size cannot be changed after writing data")
	}
	if size == 0 {
		s.variant = 256
		return s, nil
	}
	if size != 224 && size != 256 {
		return s, errors.New("sha2: invalid size, supported 224 or 256")
	}
	s.variant = size
	return s, nil
}

func (s *SHA2Basic) initHasherIfNeeded() {
	if s.hasher != nil {
		return
	}
	if s.variant == 224 {
		s.hasher = gosha256.New224()
	} else {
		s.hasher = gosha256.New()
	}
}

func (s *SHA2Basic) Write(p []byte) (n int, err error) {
	s.initHasherIfNeeded()
	return s.hasher.Write(p)
}

func (s *SHA2Basic) WriteFrom(r io.Reader) (int64, error) {
	s.initHasherIfNeeded()
	return io.Copy(s.hasher, r)
}

func (s *SHA2Basic) Reset() *SHA2Basic {
	if s.hasher != nil {
		s.hasher.Reset()
	}
	s.data = nil
	return s
}

func (s *SHA2Basic) Sum256(b []byte) *SHA2Basic {
	sum := gosha256.Sum256(b)
	return &SHA2Basic{data: sum[:], variant: 256}
}

func (s *SHA2Basic) Sum224(b []byte) *SHA2Basic {
	sum := gosha256.Sum224(b)
	return &SHA2Basic{data: sum[:], variant: 224}
}

func (s *SHA2Basic) SumStream() *SHA2Basic {
	if s.hasher == nil {
		return &SHA2Basic{}
	}
	sum := s.hasher.Sum(nil)
	return &SHA2Basic{data: sum, variant: s.variant}
}

func (s *SHA2Basic) ToBase64() string {
	return base64.StdEncoding.EncodeToString(s.data)
}

func (s *SHA2Basic) ToHex() string {
	return hex.EncodeToString(s.data)
}

func New() *SHA2Basic {
	return &SHA2Basic{variant: 256}
}
