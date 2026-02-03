package password

import (
	"crypto/rand"
	"math/big"
)

var (
	Uppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Lowercase = []byte("abcdefghijklmnopqrstuvwxyz")
	Number    = []byte("0123456789")
	Symbols   = []byte("!@#-_")
)

type PasswordBasic struct{}

func New() *PasswordBasic {
	return &PasswordBasic{}
}

func (p *PasswordBasic) pick(set []byte) byte {
	if len(set) == 0 {
		panic("empty charset")
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(set))))
	if err != nil {
		panic(err)
	}
	return set[n.Int64()]
}

func (p *PasswordBasic) shuffle(buf []byte) {
	for i := len(buf) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			panic(err)
		}
		buf[i], buf[j.Int64()] = buf[j.Int64()], buf[i]
	}
}

func (p *PasswordBasic) Generate(
	hasUpper bool,
	hasLower bool,
	hasNumber bool,
	hasSymbol bool,
	length int,
) string {
	var pool []byte
	password := make([]byte, 0, length)

	if hasUpper {
		pool = append(pool, Uppercase...)
		password = append(password, p.pick(Uppercase))
	}
	if hasLower {
		pool = append(pool, Lowercase...)
		password = append(password, p.pick(Lowercase))
	}
	if hasNumber {
		pool = append(pool, Number...)
		password = append(password, p.pick(Number))
	}
	if hasSymbol {
		pool = append(pool, Symbols...)
		password = append(password, p.pick(Symbols))
	}

	if len(pool) == 0 || length < len(password) {
		return ""
	}

	for len(password) < length {
		password = append(password, p.pick(pool))
	}

	p.shuffle(password)
	return string(password)
}
