package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Transparency struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Domain    string        `bson:"domain"`
	CreatedAt time.Time     `bson:"created_at"`
}
