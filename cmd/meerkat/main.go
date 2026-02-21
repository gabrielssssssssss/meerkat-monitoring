package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gabrielssssssssss/meerkat-monitoring/config"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/repository"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/service"
	"github.com/gabrielssssssssss/meerkat-monitoring/pkg/githarvest"
)

func main() {
	cfg := &config.Config{}
	cfg.Load("env.yaml")

	db, err := config.NewMongoDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	hitRepo := repository.NewHitRepository(db, cfg)
	transparencyRepo := repository.NewTransparencyRepository(db, cfg)

	hitService := service.NewHitService(hitRepo)
	transparencyService := service.NewTransparencyService(transparencyRepo)

	err = transparencyService.CreateDomainIndex()
	if err != nil {
		log.Fatal(err)
	}

	hit := models.Hit{
		URL:        "https://test.com",
		Path:       "/.git/config",
		Token:      "aadzadzaa",
		UserGithub: githarvest.UserGithub{},
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	}

	hitService.Create(&hit)

	resultsn, err := hitService.FindsByURL("https://test.com")
	if err != nil {
		log.Fatal(err)
	}

	for _, res := range resultsn {
		fmt.Println("Token", res.Token)
	}

}
