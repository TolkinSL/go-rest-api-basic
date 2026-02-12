package main

// CONFIG_PATH=./config/local.yaml go run ./...
// или загрузка .env через _ = godotenv.Load()
import (
	"fmt"
	"log/slog"
	"os"

	"github.com/TolkinSL/go-rest-api-basic/internal/config"
	"github.com/TolkinSL/go-rest-api-basic/internal/lib/logger/sl"
	"github.com/TolkinSL/go-rest-api-basic/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("%#v\n", cfg)

	log :=setupLogger(cfg.Env)
	log.Info("starting url-shortner", slog.String("env", cfg.Env))
	log.Debug("debug message enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}