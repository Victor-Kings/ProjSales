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
	Create(content interface{}) error
	ListById(id string) (interface{}, error)
	Update(id string, content interface{}) error
	//DeleteById(id string, collectionName string) error
}
type ExternalModel interface {
	UpdateQuery(content interface{}) primitive.D
}

type GenericRepoImp struct {
	model          ExternalModel
	collectionName string
}

func NewGenericRepo(externalModel ExternalModel, collectionName string) GenericRepo {
	return &GenericRepoImp{
		model:          externalModel,
		collectionName: collectionName,
	}
}

func getCollection(collectionName string) (*mongo.Collection, error) {
	client, err := db.GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database((db.DB)).Collection(collectionName)
	return collection, nil
}

func (g *GenericRepoImp) Create(content interface{}) error {
	collection, err := getCollection(g.collectionName)
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.TODO(), content)
	if err != nil {
		return err
	}
	return nil
}

func (g *GenericRepoImp) ListById(id string) (interface{}, error) {

	collection, err := getCollection(g.collectionName)
	if err != nil {
		return nil, err
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	content := g.model

	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: ID}}).Decode(content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (g *GenericRepoImp) Update(id string, content interface{}) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: ID}}

	query := g.model.UpdateQuery(content)

	collection, err := getCollection(g.collectionName)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(context.TODO(), filter, query)

	if err != nil {
		return err
	}

	return nil
}

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
