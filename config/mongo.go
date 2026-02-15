package config

import (
	"fmt"

	"github.com/gabrielssssssssss/meerkat-monitoring/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoDatabase(cfg *models.Config) (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%s",
		cfg.Database.Host,
		cfg.Database.Port,
	)

	client, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	return client, nil
}
