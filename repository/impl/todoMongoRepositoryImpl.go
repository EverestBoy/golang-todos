/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"go.mongodb.org/mongo-driver/bson"
	"todoGo/database"
	"todoGo/model"
	"todoGo/repository"
)

/*
@author everestboy
*/

// getting the db
var (
	dbConnection = database.NewMongoDbConnection()
)

type repo struct {
}

// New post repository
func NewFirestoreRepositry() repository.TODORepository {
	return &repo{}
}

func (r repo) FindAll() ([]model.TodoModel, error) {
	ctx, client, collection, err, cancel := dbConnection.GetCollection("test", "todo")
	defer cancel()
	defer client.Disconnect(ctx)
	if err != nil {
		println(err)
		return nil, err
	}
	//collection.InsertOne(ctx, model.TodoModel{Id: primitive.NewObjectID(), Title: "title 1", Description: "Description 1", CreatedAt: time.Now(), UpdatedAt: time.Now()})
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var todos []model.TodoModel
	println("Going throug cursor")
	for cur.Next(ctx) {
		println("inside cursor")
		var result model.TodoModel
		err := cur.Decode(&result)
		if err == nil {
			todos = append(todos, result)
		} else {
			println(err)
		}
	}
	return todos, nil
}
