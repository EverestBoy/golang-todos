/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package service

import (
	"todoGo/model"
	"todoGo/repository/impl"
)

/*
@author everestboy
*/

// getting the repo
var (
	repo = impl.NewFirestoreRepositry()
)

type TodoService interface {
	FindAllTodos() ([]model.TodoModel, error)
}

type service struct{}

func NewTodoService() TodoService {
	return &service{}
}

func (*service) FindAllTodos() ([]model.TodoModel, error) {
	return repo.FindAll()
}
