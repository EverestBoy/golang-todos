/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package service

import "todoGo/model"

/*
@author everestboy
*/

type ValidationService interface {
	ValidateTodo(todo *model.TodoModel) error
	ValidateUserRegister(user *model.User) error
	ValidateCredentials(userCredentials *model.Credential) error
}
