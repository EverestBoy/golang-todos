/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"errors"
	"todoGo/model"
	"todoGo/service"
)

type validationService struct{}

func NewValidationService() service.ValidationService {
	return &validationService{}
}

func (v validationService) ValidateTodo(todo *model.TodoModel) error {
	if todo == nil {
		err := errors.New("Post is empty")
		return err
	}
	if todo.Title == "" {
		err := errors.New("Todo Title is empty")
		return err
	}
	if todo.Description == "" {
		err := errors.New("Todo Description is empty")
		return err
	}
	return nil
}
