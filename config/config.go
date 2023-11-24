package config

import (
	"log"

	"github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig AppConfig
	DBConfig  DBConfig
}

type AppConfig struct {
	AppHost      string `env:"app_host"`
	Port         string `env:"port,required=true"`
	WriteTimeout uint32 `env:"write_timeout,required=true"`
	ReadTimeout  uint32 `env:"read_timeout,required=true"`
	IdleTimeout  uint32 `env:"idle_timeout,required=true"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&cfg)

	if err != nil {
		log.Panic(err)
	}

}

func GetConfig() Config {
	return cfg
}

type DBConfig struct {
	DB_DSN string `env:"DB_DSN,required=true"`
}
