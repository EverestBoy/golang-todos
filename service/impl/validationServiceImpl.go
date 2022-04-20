/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"errors"
	"net/mail"
	"todoGo/model"
	"todoGo/service"
)

type validationService struct{}

func NewValidationService() service.ValidationService {
	return &validationService{}
}

func (v validationService) ValidateTodo(todo *model.TodoModel) error {
	var err error = nil
	if todo == nil {
		err = errors.New("post is empty")
	}
	if todo.Title == "" {
		err = errors.New("podo Title is empty")
	}
	if todo.Description == "" {
		err = errors.New("todo Description is empty")
	}
	return err
}

func (v validationService) ValidateUserRegister(user *model.User) error {
	var err error = nil
	if user == nil {
		err = errors.New("user is empty")
	}
	_, error := mail.ParseAddress(user.Email)
	if error != nil {
		err = errors.New("email is not valid")
	}
	if user.Password == "" {
		err = errors.New("password is empty")
	}

	if user.Phone == "" {
		err = errors.New("phone is empty")
	}
	if user.Address == "" {
		err = errors.New("address is empty")
	}
	if user.Username == "" {
		err = errors.New("username is empty")
	}
	return err
}

func (v validationService) ValidateCredentials(userCredentials *model.Credential) error {
	var err error = nil
	if userCredentials.Email == "" && userCredentials.Username == "" {
		err = errors.New("username is empty")
	} else if userCredentials.Email != "" && userCredentials.Username != "" {
		err = errors.New("both email and username is not allowed")
	} else if userCredentials.Email != "" {
		_, error := mail.ParseAddress(userCredentials.Email)
		if error != nil {
			err = errors.New("email is not valid")
		}
	} else if len(userCredentials.Username) < 3 {
		err = errors.New("username is not valid")
	}
	if len(userCredentials.Password) < 5 {
		err = errors.New("password is not valid")
	}
	return err
}
