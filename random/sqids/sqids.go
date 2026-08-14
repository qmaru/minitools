package sqids

import (
	gosqids "github.com/sqids/sqids-go"
)

type Alphabet string

func (a Alphabet) String() string {
	return string(a)
}

type SqidsOptions struct {
	MinLength uint8
	Alphabet  Alphabet
	Blocklist []string
}

type SqidsBasic struct{}

const (
	AlphabetDefault  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	AlphabetReadable = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	AlphabetUpper    = "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	AlphabetLower    = "23456789abcdefghijkmnopqrstuvwxyz"
	AlphabetURLSafe  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
)

func (s *SqidsBasic) New(options SqidsOptions) (*gosqids.Sqids, error) {
	return gosqids.New(gosqids.Options{
		MinLength: options.MinLength,
		Alphabet:  options.Alphabet.String(),
		Blocklist: options.Blocklist,
	})
}

func New() *SqidsBasic {
	return new(SqidsBasic)
}
