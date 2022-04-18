/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/*
@author everestboy
*/
type MongoInterface interface {
	GetCollection(databaseName string, collectionName string) (context.Context, *mongo.Client, *mongo.Collection, error, context.CancelFunc)
}

type mongoDb struct{}

func NewMongoDbConnection() MongoInterface {
	return &mongoDb{}
}

func (*mongoDb) GetCollection(databaseName string, collectionName string) (context.Context, *mongo.Client, *mongo.Collection, error, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongo connection"))
	if err != nil {
		println("Error occured")
		println(err)
	} else {
		println("No error")
	}
	collection := client.Database(databaseName).Collection(collectionName)
	return ctx, client, collection, err, cancel
}
