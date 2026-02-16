package repository

import (
	"github.com/gabrielssssssssss/meerkat-monitoring/config"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (r transparencyRepositoryImpl) CreateDomainIndex(transparency *models.Hit) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	db := r.client.Database(r.config.Database.Name)
	collection := db.Collection("transparency")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "domain", Value: 1}},
		Options: options.Index().SetName("idx_domain_fast"),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	return nil
}
