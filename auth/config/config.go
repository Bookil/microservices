package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var ProjectRootPath = ConfigsDirPath() + "/../"

type Env int

const (
	Development Env = iota
	Production
	Test
)

var (
	CurrentEnv Env = Development
	filename   string
)

type (
	Config struct {
		ServiceName string `koanf:"service_name"`
		Mysql       Mysql  `koanf:"mysql"`
		Server      Server `koanf:"server"`
		Redis       Redis  `koanf:"redis"`
		JWT JWT 
	}

	Server struct {
		Host string `koanf:"host"`
		Port int    `koanf:"port"`
	}

	Mysql struct {
		Host     string `koanf:"host"`
		Username string `koanf:"username"`
		Password string `koanf:"password"`
		Port     int    `koanf:"port"`
		DBName   string `koanf:"db_name"`
	}

	Redis struct {
		Host     string `koanf:"host"`
		Username string `koanf:"username"`
		Password string `koanf:"password"`
		Port     int    `koanf:"port"`
		DB       int    `koanf:"db"`
	}

	JWT struct{
		SecretKey string
	}
)

func ConfigsDirPath() string {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error in generating env dir")
	}

	return filepath.Dir(f)
}

func Read() *Config {
	env := strings.ToLower(os.Getenv("AUTH_ENV"))

	log.Println("AUTH_ENV:", env)

	if len(strings.TrimSpace(env)) == 0 || env == "development" {
		CurrentEnv = Development
		filename = "config.development.yml"
	} else if env == "production" {
		CurrentEnv = Production
		filename = "config.production.yml"
	} else if env == "test" {
		CurrentEnv = Test
		filename = "config.test.yml"
	} else {
		panic(errors.New("Invalid env value set for variable AUTH_ENV: " + env))
	}

	k := koanf.New(ConfigsDirPath())
	if err := k.Load(file.Provider(fmt.Sprintf("%s/%s", ConfigsDirPath(), filename)), yaml.Parser()); err != nil {
		panic(fmt.Errorf("error loading config: %v", err))
	}
	config := &Config{}
	if err := k.Unmarshal("", config); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	secretData, secretErr := os.ReadFile(ConfigsDirPath() + "/jwt_secret.pem")
	if secretErr != nil {
		panic(secretErr)
	}

	config.JWT.SecretKey = strings.TrimSpace(string(secretData))

	log.Println("Config: ", config)
	return config
}
