package repository

import (
	"github.com/gabrielssssssssss/meerkat-monitoring/config"
	"github.com/gabrielssssssssss/meerkat-monitoring/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r hitRepositoryImpl) Create(hits *models.Hit) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	db := r.client.Database(r.config.Database.Name)
	collection := db.Collection("hits")

	_, err := collection.InsertOne(ctx, hits)
	if err != nil {
		return err
	}

	return nil
}

func (r hitRepositoryImpl) FindByToken(token string) (*models.Hit, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	db := r.client.Database(r.config.Database.Name)
	collection := db.Collection("hits")

	var results models.Hit

	err := collection.FindOne(ctx, bson.M{"token": token}).Decode(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (r hitRepositoryImpl) FindsByURL(url string) ([]models.Hit, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	db := r.client.Database(r.config.Database.Name)
	collection := db.Collection("hits")

	var results []models.Hit

	cursor, err := collection.Find(ctx, bson.M{"url": url})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var result models.Hit

		err = cursor.Decode(&result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}

	return results, nil
}
