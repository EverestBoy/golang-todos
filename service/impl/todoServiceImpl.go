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

func (s *todoService) FindTodo(id *string) (todoModel *model.TodoModel, err error) {
	return repo.FindTodo(id)
}
func (*todoService) FindAllTodos() ([]model.TodoModel, error) {
	return repo.FindAllTodo()
}

func (s *todoService) CreateTodo(todo *model.TodoModel) (*model.TodoModel, error) {
	todo.Id = primitive.NewObjectID()
	var currentTime = time.Now()
	todo.CreatedAt = currentTime
	todo.UpdatedAt = currentTime
	return repo.SaveTodo(todo)
}

func (s *todoService) UpdateTodo(id *string, todo *model.TodoModel) (*model.TodoModel, error) {
	var currentTime = time.Now()
	todo.UpdatedAt = currentTime
	return repo.UpdateTodo(id, todo)
}

func (s *todoService) DeleteTodo(id *string) (err error) {
	return repo.DeleteTodo(id)
}
