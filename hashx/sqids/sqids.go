package sqids

import (
	gosqids "github.com/sqids/sqids-go"
)

type SqidsOptions struct {
	MinLength uint8
	Alphabet  string
	Blocklist []string
}

type SqidsBasic struct{}

func (s *SqidsBasic) New(options SqidsOptions) (*gosqids.Sqids, error) {
	return gosqids.New(gosqids.Options{
		MinLength: options.MinLength,
		Alphabet:  options.Alphabet,
		Blocklist: options.Blocklist,
	})
}

func New() *SqidsBasic {
	return new(SqidsBasic)
}
