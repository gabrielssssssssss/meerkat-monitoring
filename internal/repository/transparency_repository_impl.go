package repository

import (
	"github.com/gabrielssssssssss/meerkat-monitoring/config"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TransparencyRepository interface {
	CreateDomainIndex() error
	Create(*models.Transparency) error
	FindByDomain(string) (*models.Transparency, error)
}

type transparencyRepositoryImpl struct {
	client *mongo.Client
	config *config.Config
}

func NewTransparencyRepository(client *mongo.Client, config *config.Config) TransparencyRepository {
	return &transparencyRepositoryImpl{
		client: client,
		config: config,
	}
}
