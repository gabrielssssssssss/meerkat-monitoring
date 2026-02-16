package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Transparency struct {
	ID     bson.ObjectID `bson:"_id,omitempty"`
	Domain string        `bson:"domain"`
	// UpdatedAt time.Time     `bson:"updated_at"`
	// CreatedAt time.Time     `bson:"created_at"`
}
