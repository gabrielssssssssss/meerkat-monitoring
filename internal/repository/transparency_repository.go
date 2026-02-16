package repository

import (
	"github.com/gabrielssssssssss/meerkat-monitoring/config"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (r transparencyRepositoryImpl) CreateDomainIndex() error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	db := r.client.Database(r.config.Database.Name)
	collection := db.Collection("transparency")

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"domain": 1},
		Options: options.Index().SetName("idx_domain_fast"),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r transparencyRepositoryImpl) Create(transparency *models.Transparency) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	db := r.client.Database(r.config.Database.Name)
	collection := db.Collection("transparency")

	_, err := collection.InsertOne(ctx, transparency)
	if err != nil {
		return err
	}

	return nil
}

func (r transparencyRepositoryImpl) FindByDomain(domain string) (*models.Transparency, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	db := r.client.Database(r.config.Database.Name)
	collection := db.Collection("transparency")

	var results models.Transparency

	err := collection.FindOne(ctx, bson.M{"domain": domain}).Decode(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
