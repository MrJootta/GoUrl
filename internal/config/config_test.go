package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartConfigs(t *testing.T) {
	os.Setenv("SERVER_HOST", "localhost")
	os.Setenv("SERVER_PORT", "4040")

	os.Setenv("DB_DATABASE", "shortlink")
	os.Setenv("DB_HOST", "56.56.56.56")
	os.Setenv("DB_PASSWORD", "testing")
	os.Setenv("DB_PORT", "5642")
	os.Setenv("DB_USERNAME", "testing")

	t.Run("get server configs", func(t *testing.T) {
		configs := StartConfigs()

		assert.NotEmpty(t, configs)
		assert.Equal(t, "localhost", configs.ServerHost)
		assert.Equal(t, "4040", configs.ServerPort)
	})

	t.Run("get database configs", func(t *testing.T) {
		configs := StartConfigs()

		assert.NotEmpty(t, configs)
		assert.Equal(t, "shortlink", configs.DBDatabase)
		assert.Equal(t, "56.56.56.56", configs.DBHost)
		assert.Equal(t, "testing", configs.DBPass)
		assert.Equal(t, "5642", configs.DBPort)
		assert.Equal(t, "testing", configs.DBUser)
	})

}
