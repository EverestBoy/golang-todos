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
	FindTodo(todoId *string, userId *string) (*model.TodoModel, error)
	FindAllTodo(userId *string) ([]model.TodoModel, error)
	SaveTodo(todo *model.TodoModel, userId *string) (*model.TodoModel, error)
	UpdateTodo(id *string, userId *string, todo *model.TodoModel) (*model.TodoModel, error)
	DeleteTodo(id *string, userId *string) error
}
