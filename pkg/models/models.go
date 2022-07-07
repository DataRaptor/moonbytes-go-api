package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AlphaUser struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Uuid          string             `json:"uuid" bson:"uuid,omitempty"`
	PublicKey     string             `json:"public_key" bson:"public_key,omitempty"`
	ApiKey        string             `json:"api_key" bson:"api_key,omitempty"`
	RequestsCount int                `json:"requests_count" bson:"requests_count,omitempty"`
}
