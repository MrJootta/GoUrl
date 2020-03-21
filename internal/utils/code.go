package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
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

func URLBuilder(code string) string {
	return fmt.Sprintf("http://%s:%s/%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"), code)
}
