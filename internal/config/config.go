package config

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/caarlos0/env/v6"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type сonfig struct {
	ServerPort     ServerConfig
	DatabaseConfig DBConfig
}

type ServerConfig struct {
	Port string `env:"SERVER_PORT" envDefault:":8080"`
}

type DBConfig struct {
	URL string `env:"DB_URL" envDefault:"postgresql://postgres:postgres@localhost:5432/postgres"` // цепляемся к дефолтному Postgres
}

var (
	config *сonfig //nolint: gochecknoglobals
)

func GetConfig() *сonfig {
	var err error

	config, err = newConfig()
	if err != nil {
		log.Fatalf("new instance config: %v", err)
	}

	return config
}

func newConfig() (*сonfig, error) {
	conf := &сonfig{}

	err := godotenv.Load()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return conf, fmt.Errorf("load config: %w", err)
		}
	}

	err = env.ParseWithFuncs(conf, map[reflect.Type]env.ParserFunc{
		reflect.TypeOf(rsa.PublicKey{}): func(v string) (interface{}, error) {
			return base64.StdEncoding.DecodeString(v)
		},
	})
	if err != nil {
		return nil, fmt.Errorf("parse config from env: %w", err)
	}

	return conf, nil
}
