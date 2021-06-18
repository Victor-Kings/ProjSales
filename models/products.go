package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// import (
// 	"context"
// 	"log"

// 	"github.com/Victor-Kings/ProjSales/db"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// const productsCollection = "products"

// type ProductsDB interface {
// 	CreateProducts(product Products) error
// 	ListAllProducts() ([]Products, error)
// 	ListByIdProducts(id string) (*Products, error)
// 	UpdateProduct(id string, product Products) error
// 	DeleteProductById(id string) error
// }

type ProductsDBImp struct{}

type Products struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Mark   string             `json:"mark,,omitempty" bson:"mark,,omitempty"`
	Amount int                `json:"amount,,omitempty" bson:"amount,,omitempty"`
	Model  string             `json:"model" bson:"model"`
	Color  string             `json:"color" bson:"color"`
	Price  float64            `json:"price" bson:"price"`
}

// func NewProductsDB() ProductsDB {
// 	return &ProductsDBImp{}
// }

// func getCollectionMongo() (*mongo.Collection, error) {
// 	client, err := db.GetMongoClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	collection := client.Database((db.DB)).Collection(productsCollection)
// 	return collection, nil
// }

// func (p ProductsDBImp) CreateProducts(product Products) error {
// 	collection, err := getCollectionMongo()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = collection.InsertOne(context.TODO(), product)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (p ProductsDBImp) ListAllProducts() ([]Products, error) {
// 	collection, err := getCollectionMongo()
// 	if err != nil {
// 		return nil, err
// 	}

// 	cursor, err2 := collection.Find(context.TODO(), bson.D{})
// 	if err2 != nil {
// 		return nil, err2
// 	}
// 	products := []Products{}

// 	if err = cursor.All(context.TODO(), &products); err != nil {
// 		log.Fatal(err)
// 	}

// 	return products, nil
// }

// func (p ProductsDBImp) ListByIdProducts(id string) (*Products, error) {

// 	collection, err := getCollectionMongo()
// 	if err != nil {
// 		return nil, err
// 	}

// 	ID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	product := Products{}

// 	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: ID}}).Decode(&product)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &product, nil
// }

// func (p ProductsDBImp) UpdateProduct(id string, product Products) error {
// 	ID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	filter := bson.D{primitive.E{Key: "_id", Value: ID}}

// 	query := bson.D{
// 		primitive.E{
// 			Key: "$set",
// 			Value: bson.D{
// 				primitive.E{
// 					Key:   "mark",
// 					Value: product.Mark,
// 				},
// 				primitive.E{
// 					Key:   "amount",
// 					Value: product.Amount,
// 				},
// 				primitive.E{
// 					Key:   "model",
// 					Value: product.Model,
// 				},
// 				primitive.E{
// 					Key:   "color",
// 					Value: product.Color,
// 				},
// 				primitive.E{
// 					Key:   "price",
// 					Value: product.Price,
// 				},
// 			},
// 		},
// 	}

// 	collection, err := getCollectionMongo()
// 	if err != nil {
// 		return err
// 	}

// 	_, err = collection.UpdateOne(context.TODO(), filter, query)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p ProductsDBImp) DeleteProductById(id string) error {

// 	collection, err := getCollectionMongo()
// 	if err != nil {
// 		return err
// 	}

// 	ID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: ID}})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
