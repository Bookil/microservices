package config_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/Bookil/microservices/user/config"
	"github.com/stretchr/testify/require"
)

func TestDevelopmentConfig(t *testing.T) {
	err := checkAndSet("development")
	if err != nil {
		if isFileNotExists(err) {
			return
		}
		require.Error(t, err)
	}

	configs := config.Read()
	require.NotEmpty(t, configs)
	require.Equal(t, config.CurrentEnv, config.Development)
}

func TestProductionConfig(t *testing.T) {
	err := checkAndSet("production")
	if err != nil {
		if isFileNotExists(err) {
			return
		}
		require.Error(t, err)
	}

	configs := config.Read()
	require.NotEmpty(t, configs)
	require.Equal(t, config.CurrentEnv, config.Production)
}

func TestTestConfig(t *testing.T) {
	err := checkAndSet("test")
	if err != nil {
		if isFileNotExists(err) {
			return
		}
		require.Error(t, err)
	}

	configs := config.Read()
	require.NotEmpty(t, configs)
	require.Equal(t, config.CurrentEnv, config.Test)
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

func checkAndSet(env string) error {
	_, err := os.Stat(fmt.Sprintf("config.%s.yml", env))
	if err != nil {
		return err
	}

	os.Setenv("USER_ENV", env)

	return nil
}

func isFileNotExists(err error) bool {
	isFileNotExists := errors.Unwrap(err).Error() == "no such file or directory"

	return isFileNotExists
}
