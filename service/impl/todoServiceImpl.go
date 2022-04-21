/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"todoGo/model"
	"todoGo/repository/impl"
	"todoGo/service"
)

// getting the repo
var (
	repo = impl.NewTodoRepository()
)

type todoService struct{}

func NewTodoService() service.TodoService {
	return &todoService{}
}

func (s *todoService) FindTodo(id *string, userId *string) (todoModel *model.TodoModel, err error) {
	return repo.FindTodo(id, userId)
}
func (*todoService) FindAllTodos(userId *string) ([]model.TodoModel, error) {
	return repo.FindAllTodo(userId)
}

func (s *todoService) CreateTodo(todo *model.TodoModel, userId *string) (*model.TodoModel, error) {
	todo.Id = primitive.NewObjectID()
	var currentTime = time.Now()
	todo.CreatedAt = currentTime
	todo.UpdatedAt = currentTime
	todo.UserId = *userId
	return repo.SaveTodo(todo, userId)
}

func (s *todoService) UpdateTodo(id *string, userId *string, todo *model.TodoModel) (*model.TodoModel, error) {
	var currentTime = time.Now()
	todo.UpdatedAt = currentTime
	return repo.UpdateTodo(id, userId, todo)
}

func (s *todoService) DeleteTodo(id *string, userId *string) (err error) {
	return repo.DeleteTodo(id, userId)
}
