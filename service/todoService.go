/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package service

import (
	"todoGo/model"
)

/*
@author everestboy
*/

type TodoService interface {
	FindTodo(id *string) (todoModel *model.TodoModel, err error)
	FindAllTodos() ([]model.TodoModel, error)
	CreateTodo(todo *model.TodoModel) (*model.TodoModel, error)
	UpdateTodo(id *string, todo *model.TodoModel) (*model.TodoModel, error)
	DeleteTodo(id *string) (err error)
}
