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
	AlphabetDefault  Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	AlphabetReadable Alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	AlphabetUpper    Alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	AlphabetLower    Alphabet = "23456789abcdefghijkmnopqrstuvwxyz"
	AlphabetURLSafe  Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
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
