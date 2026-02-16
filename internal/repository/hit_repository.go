package repository

import (
	"github.com/gabrielssssssssss/meerkat-monitoring/config"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type HitRepository interface {
	Create(*models.Hit) error
	FindByToken(string) (*models.Hit, error)
	FindsByURL(string) ([]models.Hit, error)
}

type hitRepositoryImpl struct {
	client *mongo.Client
	config *config.Config
}

func NewHitRepository(client *mongo.Client, config *config.Config) HitRepository {
	return &hitRepositoryImpl{
		client: client,
		config: config,
	}
}
