package models

import (
	"context"
	"log"

	"github.com/Victor-Kings/ProjSales/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const productsCollection = "products"

type ProductsDB interface {
	CreateProducts(product Products) error
	ListAllProducts() ([]Products, error)
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

func (p ProductsDBImp) ListAllProducts() ([]Products, error) {
	client, err := db.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database((db.DB)).Collection(productsCollection)

	cursor, err2 := collection.Find(context.TODO(), bson.D{})
	if err2 != nil {
		return nil, err2
	}
	products := []Products{}

	if err = cursor.All(context.TODO(), &products); err != nil {
		log.Fatal(err)
	}

	return products, nil
}
