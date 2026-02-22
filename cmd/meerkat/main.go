package main

import (
	"os"

	"github.com/gabrielssssssssss/meerkat-monitoring/config"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/repository"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/runner"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/service"
	"github.com/gabrielssssssssss/meerkat-monitoring/pkg/githarvest"
	"github.com/gabrielssssssssss/meerkat-monitoring/pkg/transparency"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func main() {
	options := runner.ParseOptions()

	cfg := &config.Config{}
	cfg.Load(options.Cfg)

	db, err := config.NewMongoDatabase(cfg)
	if err != nil {
		log.Error().
			Err(err).
			Str("component", "config.NewMongoDatabase").
			Msg("Connection failed to MongoDB")
	}

	hitRepo := repository.NewHitRepository(db, cfg)
	transparencyRepo := repository.NewTransparencyRepository(db, cfg)

	hitService := service.NewHitService(hitRepo)
	transparencyService := service.NewTransparencyService(transparencyRepo)

	gitHarvestClient := githarvest.NewClient()
	transparencyClient := transparency.NewClient()

	runner := runner.NewRunner(options, cfg, hitService, transparencyService, gitHarvestClient, transparencyClient)

	runner.RunScanner()
}
