package sqids

import (
	"github.com/sqids/sqids-go"
)

type SqidsOptions struct {
	MinLength uint8
	Alphabet  string
	Blocklist []string
}

type SqidsBasic struct{}

func (s *SqidsBasic) New(options SqidsOptions) (*sqids.Sqids, error) {
	return sqids.New(sqids.Options{
		MinLength: options.MinLength,
		Alphabet:  options.Alphabet,
		Blocklist: options.Blocklist,
	})
}

func New() *SqidsBasic {
	return new(SqidsBasic)
}
