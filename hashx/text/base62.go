package text

import (
	"fmt"
)

var (
	base62Charset = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
)

func (t *TextBasic) Base62Encode(s []byte) string {
	if len(s) == 0 {
		return ""
	}

	data := make([]byte, len(s))
	copy(data, s)

	var encoded []byte
	for len(data) > 0 {
		var remainder int
		var next []byte
		for _, b := range data {
			acc := int(b) + remainder*256
			quo := acc / 62
			remainder = acc % 62
			if len(next) > 0 || quo > 0 {
				next = append(next, byte(quo))
			}
		}
		encoded = append(encoded, base62Charset[remainder])
		data = next
	}

	for i, j := 0, len(encoded)-1; i < j; i, j = i+1, j-1 {
		encoded[i], encoded[j] = encoded[j], encoded[i]
	}

	return string(encoded)
}

func (t *TextBasic) Base62Decoding(s string) (*TextBasic, error) {
	if s == "" {
		return nil, fmt.Errorf("input is required")
	}

	base62Index := map[byte]int{}
	for i, b := range base62Charset {
		base62Index[b] = i
	}

	result := []byte{0}
	for i := range len(s) {
		val, ok := base62Index[s[i]]
		if !ok {
			return nil, fmt.Errorf("invalid base62 character: %c", s[i])
		}

		carry := val
		for j := range result {
			carry += int(result[j]) * 62
			result[j] = byte(carry % 256)
			carry /= 256
		}

		for carry > 0 {
			result = append(result, byte(carry%256))
			carry /= 256
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return &TextBasic{data: result}, nil
}
