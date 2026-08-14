package nanoid

import (
	"github.com/matoous/go-nanoid/v2"
)

type Alphabet string

func (a Alphabet) String() string {
	return string(a)
}

const (
	AlphabetDefault  Alphabet = "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphabetBase62   Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	AlphabetReadable Alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	AlphabetUpper    Alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	AlphabetLower    Alphabet = "23456789abcdefghijkmnopqrstuvwxyz"
)

type NanoidBasic struct{}

func (s *NanoidBasic) New(l ...int) (string, error) {
	return gonanoid.New(l...)
}

func (s *NanoidBasic) Generate(alphabet Alphabet, size int) (string, error) {
	return gonanoid.Generate(alphabet.String(), size)
}

func New() *NanoidBasic {
	return new(NanoidBasic)
}
