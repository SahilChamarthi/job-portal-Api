package config

import (
	"log"

	"github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig              AppConfig
	DBConfig               DBConfig
	RedisConfig            RedisConfig
	PrivatePublicPemConfig PrivatePublicPemConfig
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

type AppConfig struct {
	AppHost      string `env:"app_host"`
	Port         string `env:"port,required=true"`
	WriteTimeout uint32 `env:"write_timeout,required=true"`
	ReadTimeout  uint32 `env:"read_timeout,required=true"`
	IdleTimeout  uint32 `env:"idle_timeout,required=true"`
}

type DBConfig struct {
	DB_DSN string `env:"DB_DSN,required=true"`
}

type RedisConfig struct {
	Addr     string `env:"addr,required=true"`
	Password string `env:"password"`
	DB       int    `env:"db"`
}

type PrivatePublicPemConfig struct {
	PrivatePem string `env:"private_pem,required=true"`
	PublicPem  string `env:"public_pem,required=true"`
}
