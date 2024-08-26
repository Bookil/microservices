package config_test

import (
	"os"
	"testing"

	"github.com/Bookil/microservices/user/config"
	"github.com/stretchr/testify/require"
)

func TestDevelopmentConfig(t *testing.T) {
	os.Setenv("USER_ENV", "development")

	configs := config.Read()
	require.NotEmpty(t, configs)
	require.Equal(t, config.CurrentEnv, config.Development)
}

func TestTestConfig(t *testing.T) {
	os.Setenv("USER_ENV", "test")

	configs := config.Read()
	require.NotEmpty(t, configs)
	require.Equal(t, config.CurrentEnv, config.Test)
}

func TestProductionConfig(t *testing.T) {
	os.Setenv("USER_ENV", "production")

	defer func() {
		r := recover()
		if r == nil {
			require.Fail(t, "should fatal")
		}
	}()

	config.Read()
}

func TestInvalidEnv(t *testing.T) {
	os.Setenv("USER_ENV", "invalid")

	defer func() {
		r := recover()
		if r == nil {
			require.Fail(t, "should panic")
		}
	}()

	config.Read()
}
