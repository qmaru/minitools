package nanoid

import (
	"github.com/matoous/go-nanoid/v2"
)

type NanoidBasic struct{}

func (s *NanoidBasic) New(l ...int) (string, error) {
	return gonanoid.New()
}

func (s *NanoidBasic) Generate(alphabet string, size int) (string, error) {
	return gonanoid.Generate(alphabet, size)
}

func New() *NanoidBasic {
	return new(NanoidBasic)
}
