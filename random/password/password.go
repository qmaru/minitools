package password

import (
	"crypto/rand"
	"math/big"
)

type Charset struct {
	Uppercase string
	Lowercase string
	Number    string
	Symbols   string
}

func DefaultCharset() Charset {
	return Charset{
		Uppercase: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Lowercase: "abcdefghijklmnopqrstuvwxyz",
		Number:    "0123456789",
		Symbols:   "!@#$%^&*()-_+=",
	}
}

type PasswordBasic struct {
	Charset Charset
}

func New() *PasswordBasic {
	return &PasswordBasic{
		Charset: DefaultCharset(),
	}
}

func NewWithCharset(charset Charset) *PasswordBasic {
	return &PasswordBasic{
		Charset: charset,
	}
}

func (p *PasswordBasic) pick(set string) byte {
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

func (p *PasswordBasic) Generate(length int) string {
	var pool string
	password := make([]byte, 0, length)

	if p.Charset.Uppercase != "" {
		if p.Charset.Uppercase == "" {
			return ""
		}

		pool += p.Charset.Uppercase
		password = append(password, p.pick(p.Charset.Uppercase))
	}
	if p.Charset.Lowercase != "" {
		if p.Charset.Lowercase == "" {
			return ""
		}

		pool += p.Charset.Lowercase
		password = append(password, p.pick(p.Charset.Lowercase))
	}
	if p.Charset.Number != "" {
		if p.Charset.Number == "" {
			return ""
		}

		pool += p.Charset.Number
		password = append(password, p.pick(p.Charset.Number))
	}
	if p.Charset.Symbols != "" {
		if p.Charset.Symbols == "" {
			return ""
		}

		pool += p.Charset.Symbols
		password = append(password, p.pick(p.Charset.Symbols))
	}

	if pool == "" || length < len(password) {
		return ""
	}

	for len(password) < length {
		password = append(password, p.pick(pool))
	}

	p.shuffle(password)
	return string(password)
}
