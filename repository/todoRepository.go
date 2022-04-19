/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package repository

import "todoGo/model"

/*
@author everestboy
*/

type TODORepository interface {
	FindTodo(todoId *string) (*model.TodoModel, error)
	FindAllTodo() ([]model.TodoModel, error)
	SaveTodo(todo *model.TodoModel) (*model.TodoModel, error)
	UpdateTodo(id *string, todo *model.TodoModel) (*model.TodoModel, error)
	DeleteTodo(id *string) error
}
