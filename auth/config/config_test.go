package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigsDirPath(t *testing.T) {
	t.Log(ConfigsDirPath())
	t.Log(ProjectRootPath)
}

func TestDevelopmentConfig(t *testing.T) {
	os.Setenv("AUTH_ENV", "development")

	configs := Read()
	require.NotEmpty(t, configs)
	require.Equal(t, CurrentEnv, Development)
}
func TestTestConfig(t *testing.T) {
	os.Setenv("AUTH_ENV", "test")

	configs := Read()
	require.NotEmpty(t, configs)
	require.Equal(t, CurrentEnv, Test)
}

func TestProductionConfig(t *testing.T) {
	os.Setenv("AUTH_ENV", "production")

	defer func() {
		r := recover()
		if r == nil {
			require.Fail(t,"should fatal")
		}
	}()

	Read()
}

func TestInvalidEnv(t *testing.T) {
	os.Setenv("AUTH_ENV", "invalid")

	defer func() {
		r := recover()
		if r == nil {
			require.Fail(t,"should panic")
		}
	}()

	Read()
}
