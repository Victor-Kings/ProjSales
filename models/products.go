package models

import (
	"context"

	"github.com/Victor-Kings/ProjSales/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const productsCollection = "products"

type ProductsDB interface {
	CreateProducts(product Products) error
}

type ProductsDBImp struct{}

type Products struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Mark   string             `json:"mark,,omitempty" bson:"mark,,omitempty"`
	Amount int                `json:"amount,,omitempty" bson:"amount,,omitempty"`
	Model  string             `json:"model" bson:"model"`
	Color  string             `json:"color" bson:"color"`
	Price  float64            `json:"price" bson:"price"`
}

func NewProductsDB() ProductsDB {
	return &ProductsDBImp{}
}

func (p ProductsDBImp) CreateProducts(product Products) error {
	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database((db.DB)).Collection(productsCollection)

	_, err = collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}
