package config

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigsDirPath(t *testing.T) {
	t.Log(ConfigsDirPath())
	t.Log(ProjectRootPath)
}

func TestDevelopmentConfig(t *testing.T) {
	err := checkAndSet("development")
	if err != nil{
		if isFileNotExists(err){
			return
		}
		require.Error(t,err)
	}

	configs := Read()
	require.NotEmpty(t, configs)
	require.Equal(t, CurrentEnv, Development)
}

func TestProductionConfig(t *testing.T) {
	err := checkAndSet("production")
	if err != nil{
		if isFileNotExists(err){
			return
		}
		require.Error(t,err)
	}

	configs := Read()
	require.NotEmpty(t, configs)
	require.Equal(t, CurrentEnv, Production)
}

func TestTestConfig(t *testing.T) {
	err := checkAndSet("test")
	if err != nil{
		if isFileNotExists(err){
			return
		}
		require.Error(t,err)
	}

	configs := Read()
	require.NotEmpty(t, configs)
	require.Equal(t, CurrentEnv, Test)
}

func TestInvalidEnv(t *testing.T) {
	os.Setenv("EMAIL_ENV", "invalid")

	defer func() {
		r := recover()
		if r == nil {
			require.Fail(t, "should panic")
		}
	}()
	Read()
}



func checkAndSet(env string) error {
	_, err := os.Stat(fmt.Sprintf("config.%s.yml", env))
	if err != nil {
		return err
	}

	os.Setenv("EMAIL_ENV", env)

	return nil
}

func isFileNotExists(err error)bool{
	isFileNotExists := errors.Unwrap(err).Error() == "no such file or directory" 

	return isFileNotExists
}