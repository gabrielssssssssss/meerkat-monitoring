package models

import (
	"time"

	"github.com/gabrielssssssssss/meerkat-monitoring/pkg/githarvest"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Hit struct {
	ID         bson.ObjectID         `bson:"_id,omitempty"`
	URL        string                `bson:"url"`
	Path       string                `bson:"path"`
	Token      string                `bson:"token"`
	UserGithub githarvest.UserGithub `bson:"user_github"`
	UpdatedAt  time.Time             `bson:"updated_at"`
	CreatedAt  time.Time             `bson:"created_at"`
}
