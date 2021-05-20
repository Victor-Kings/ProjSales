package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Products struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Mark   string             `json:"mark" bson:"mark"`
	Amount int                `json:"amount" bson:"amount"`
	Model  string             `json:"model" bson:"model"`
	Color  string             `json:"color" bson:"color"`
	Price  float64            `json:"price" bson:"price"`
}
