package utils

import (
	"crypto/rand"
	"math/big"
	"unicode"
)

const length = 6

func GenerateCode() (string, error) {
	var code string

	for len(code) < length {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}

		n := num.Int64()
		if unicode.IsLetter(rune(n)) {
			code += string(n)
		}
	}

	return code, nil
}
