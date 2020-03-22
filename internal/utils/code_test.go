package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {
	t.Run("generate code with length of 6", func(t *testing.T) {
		code, err := GenerateCode()

		assert.Nil(t, err)
		assert.Equal(t, 6, len(code))
	})
}

func TestURLBuilder(t *testing.T) {
	t.Run("get url with code", func(t *testing.T) {
		url := URLBuilder("randomCode")

		assert.Equal(t, "http://:/randomCode", url)
	})
}
