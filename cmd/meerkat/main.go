package main

import (
	"log"

	"github.com/gabrielssssssssss/meerkat-monitoring/config"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/repository"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/runner"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/service"
	"github.com/gabrielssssssssss/meerkat-monitoring/pkg/githarvest"
	"github.com/gabrielssssssssss/meerkat-monitoring/pkg/transparency"
)

func main() {
	options := runner.ParseOptions()

	cfg := &config.Config{}
	cfg.Load(options.Cfg)

	db, err := config.NewMongoDatabase(cfg)
	if err != nil {
		log.Fatal(err)
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
