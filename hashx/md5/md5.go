package md5

import (
	gomd5 "crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"io"
)

type MD5Basic struct {
	data   []byte
	hasher hash.Hash
}

func (m *MD5Basic) Write(p []byte) (n int, err error) {
	if m.hasher == nil {
		m.hasher = gomd5.New()
	}
	return m.hasher.Write(p)
}

func (m *MD5Basic) WriteFrom(r io.Reader) (int64, error) {
	if m.hasher == nil {
		m.hasher = gomd5.New()
	}
	return io.Copy(m.hasher, r)
}

func (m *MD5Basic) Reset() *MD5Basic {
	if m.hasher != nil {
		m.hasher.Reset()
	}
	m.data = nil
	return m
}

func (m *MD5Basic) Sum(b []byte) *MD5Basic {
	sum := gomd5.Sum(b)
	return &MD5Basic{data: sum[:]}
}

func (m *MD5Basic) SumStream() *MD5Basic {
	if m.hasher == nil {
		return &MD5Basic{}
	}
	sum := m.hasher.Sum(nil)
	return &MD5Basic{data: sum}
}

func (m *MD5Basic) ToBase64() string {
	return base64.StdEncoding.EncodeToString(m.data)
}

func (m *MD5Basic) ToHex() string {
	return hex.EncodeToString(m.data)
}

func New() *MD5Basic {
	return &MD5Basic{}
}
