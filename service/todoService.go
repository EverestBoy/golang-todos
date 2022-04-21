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
	FindTodo(id *string, userId *string) (todoModel *model.TodoModel, err error)
	FindAllTodos(userId *string) ([]model.TodoModel, error)
	CreateTodo(todo *model.TodoModel, userId *string) (*model.TodoModel, error)
	UpdateTodo(id *string, userId *string, todo *model.TodoModel) (*model.TodoModel, error)
	DeleteTodo(id *string, userId *string) (err error)
}
