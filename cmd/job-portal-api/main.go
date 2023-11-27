package main

import (
	"fmt"
	"net/http"
	"project/config"
	"project/internal/auth"
	"project/internal/database"
	"project/internal/handlers"
	redispack "project/internal/redisPack"
	"project/internal/repository"
	"project/internal/services"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
}
func startApp() error {
	cfg := config.GetConfig()
	log.Info().Msg("started main")
	privatePEM := cfg.PrivatePublicPemConfig.PrivatePem

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatePEM))
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}

	publicPEM := cfg.PrivatePublicPemConfig.PublicPem

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicPEM))
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("cannot create auth instance %w", err)
	}

	redisClient := redispack.NewRedisClient()
	rc := redispack.NewRedisConnection(redisClient)

	db, err := database.DataBaseConnect()
	if err != nil {
		return err
	}
	repo, err := repository.NewRepo(db)
	if err != nil {
		return err
	}

	se, err := services.NewServices(repo, rc)

	if err != nil {
		return err
	}

	api := http.Server{ //server configuration
		Addr:         fmt.Sprintf("%s:%s", cfg.AppConfig.AppHost, cfg.AppConfig.Port),
		ReadTimeout:  time.Duration(cfg.AppConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.AppConfig.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.AppConfig.IdleTimeout) * time.Second,
		Handler:      handlers.Api(a, se),
	}
	api.ListenAndServe()

	return nil

}
