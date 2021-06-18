package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// import (
// 	"context"
// 	"log"

// 	"github.com/Victor-Kings/ProjSales/db"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// const consumerCollection = "consumer"

// type ConsumerDB interface {
// 	Createconsumer(consumer Consumer) error
// 	ListAllconsumer() ([]Consumer, error)
// 	ListByIdconsumer(id string) (*Consumer, error)
// 	Updateconsumer(id string, consumer Consumer) error
// 	DeleteconsumerById(id string) error
// }

// type ConsumerDBImp struct{}

type Consumer struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,,omitempty" bson:"name,,omitempty"`
}

// func NewconsumerDB() ConsumerDB {
// 	return &ConsumerDBImp{}
// }

// func getCollectionMongoConsumer() (*mongo.Collection, error) {
// 	client, err := db.GetMongoClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	collection := client.Database((db.DB)).Collection(consumerCollection)
// 	return collection, nil
// }

// func (p ConsumerDBImp) Createconsumer(consumer Consumer) error {
// 	collection, err := getCollectionMongoConsumer()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = collection.InsertOne(context.TODO(), consumer)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (p ConsumerDBImp) ListAllconsumer() ([]Consumer, error) {
// 	collection, err := getCollectionMongoConsumer()
// 	if err != nil {
// 		return nil, err
// 	}

// 	cursor, err2 := collection.Find(context.TODO(), bson.D{})
// 	if err2 != nil {
// 		return nil, err2
// 	}
// 	consumer := []Consumer{}

// 	if err = cursor.All(context.TODO(), &consumer); err != nil {
// 		log.Fatal(err)
// 	}

// 	return consumer, nil
// }

// func (p ConsumerDBImp) ListByIdconsumer(id string) (*Consumer, error) {

// 	collection, err := getCollectionMongoConsumer()
// 	if err != nil {
// 		return nil, err
// 	}

// 	ID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	consumer := Consumer{}

// 	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: ID}}).Decode(&consumer)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &consumer, nil
// }

// func (p ConsumerDBImp) Updateconsumer(id string, consumer Consumer) error {
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
// 					Key:   "name",
// 					Value: consumer.Name,
// 				},
// 			},
// 		},
// 	}

// 	collection, err := getCollectionMongoConsumer()
// 	if err != nil {
// 		return err
// 	}

// 	_, err = collection.UpdateOne(context.TODO(), filter, query)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p ConsumerDBImp) DeleteconsumerById(id string) error {

// 	collection, err := getCollectionMongoConsumer()
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
