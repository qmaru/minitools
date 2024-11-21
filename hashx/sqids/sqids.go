package sqids

import (
	"github.com/sqids/sqids-go"
)

type SqidsBasic struct{}

func (s *SqidsBasic) New(minLen uint8) (*sqids.Sqids, error) {
	return sqids.New(sqids.Options{MinLength: minLen})
}

func New() *SqidsBasic {
	return new(SqidsBasic)
}
