/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
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

func NewTodoRepository() repository.TODORepository {
	return &repo{}
}

func (r repo) FindTodo(todoId *string) (*model.TodoModel, error) {
	err, ctx, client, collection, cancel := dbConnection.GetCollection("test", "todo")
	defer cancel()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}
	var todoResponse model.TodoModel
	id, _ := primitive.ObjectIDFromHex(*todoId)
	todoResult := collection.FindOne(ctx, bson.M{"_id": id})
	err = todoResult.Decode(&todoResponse)
	if err != nil {
		return nil, err
	}
	return &todoResponse, nil
}

func (r repo) FindAllTodo() ([]model.TodoModel, error) {
	err, ctx, client, collection, cancel := dbConnection.GetCollection("test", "todo")
	if err != nil {
		return nil, err
	}
	defer cancel()
	defer client.Disconnect(ctx)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var todos []model.TodoModel
	for cur.Next(ctx) {
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

func (r repo) SaveTodo(todo *model.TodoModel) (*model.TodoModel, error) {
	err, ctx, client, collection, cancel := dbConnection.GetCollection("test", "todo")
	defer cancel()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}
	cur, err := collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	insertedId := cur.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("Inserted todo %s", insertedId)
	insertedTodo, err := r.FindTodo(&insertedId)
	return insertedTodo, err
}

func (r repo) UpdateTodo(id *string, todo *model.TodoModel) (*model.TodoModel, error) {
	err, ctx, client, collection, cancel := dbConnection.GetCollection("test", "todo")
	defer cancel()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}
	todoModel, err := r.FindTodo(id)
	if err != nil {
		return nil, err
	}
	todoModel.Title = todo.Title
	todoModel.Description = todo.Description
	todoModel.UpdatedAt = todo.UpdatedAt

	filter := bson.M{"_id": todoModel.Id}
	update := bson.D{
		{"$set", todoModel},
	}
	updated, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	println(updated.UpsertedID)
	log.Println("Updated todo " + *id)
	return r.FindTodo(id)
}

func (r repo) DeleteTodo(id *string) error {
	err, ctx, client, collection, cancel := dbConnection.GetCollection("test", "todo")
	defer cancel()
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	if err != nil {
		return err
	}
	todoModel, err := r.FindTodo(id)
	if err != nil {
		return errors.Errorf("Cannot perform delete operation")
	}
	filter := bson.M{"_id": todoModel.Id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	log.Printf("Deleted %d records\n", result.DeletedCount)
	return nil
}
