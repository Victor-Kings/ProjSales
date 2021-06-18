package repository

import (
	"context"

	"github.com/Victor-Kings/ProjSales/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//const productsCollection = "products"

type GenericRepo interface {
	//Create(product Products, collection string) error

	ListById(a interface{}, id string, collectionName string) (interface{}, error)
	//Update(id string, product Products, collectionName string) error
	//DeleteById(id string, collectionName string) error
}

// type generic struct {
// 	mConsumer models.Consumer
// 	mProducts models.Products
// }
type GenericRepoImp struct {
}

func NewGenericRepo() GenericRepo {
	return &GenericRepoImp{}
}

func getCollection(collectionName string) (*mongo.Collection, error) {
	client, err := db.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database((db.DB)).Collection(collectionName)
	return collection, nil
}

// func (p GenericRepoImp) Create(product Products, collectionName string) error {
// 	collection, err := getCollection(collectionName)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = collection.InsertOne(context.TODO(), product)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (p GenericRepoImp) ListById(a interface{}, id, collectionName string) (interface{}, error) {

	collection, err := getCollection(collectionName)
	if err != nil {
		return nil, err
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	content := a

	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: ID}}).Decode(&content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

// func (p GenericRepoImp) Update(id string, product Products, collectionName string, query primitive.D) error {
// 	ID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	filter := bson.D{primitive.E{Key: "_id", Value: ID}}

// 	// query := bson.D{
// 	// 	primitive.E{
// 	// 		Key: "$set",
// 	// 		Value: bson.D{
// 	// 			primitive.E{
// 	// 				Key:   "mark",
// 	// 				Value: product.Mark,
// 	// 			},
// 	// 			primitive.E{
// 	// 				Key:   "amount",
// 	// 				Value: product.Amount,
// 	// 			},
// 	// 			primitive.E{
// 	// 				Key:   "Repo",
// 	// 				Value: product.Repo,
// 	// 			},
// 	// 			primitive.E{
// 	// 				Key:   "color",
// 	// 				Value: product.Color,
// 	// 			},
// 	// 			primitive.E{
// 	// 				Key:   "price",
// 	// 				Value: product.Price,
// 	// 			},
// 	// 		},
// 	// 	},
// 	// }

// 	collection, err := getCollection(collectionName)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = collection.UpdateOne(context.TODO(), filter, query)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p GenericRepoImp) DeleteById(id string, collectionName string) error {

// 	collection, err := getCollection(collectionName)
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
