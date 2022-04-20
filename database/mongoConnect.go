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
	GetCollection(databaseName string, collectionName string) (error, context.Context, *mongo.Client, *mongo.Collection, context.CancelFunc)
}

const (
	mongoURI = ""
)

type mongoDb struct{}

func NewMongoDbConnection() MongoInterface {
	return &mongoDb{}
}

func (*mongoDb) GetCollection(databaseName string, collectionName string) (error, context.Context, *mongo.Client, *mongo.Collection, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err, nil, nil, nil, nil
	}
	collection := client.Database(databaseName).Collection(collectionName)
	return nil, ctx, client, collection, cancel
}
