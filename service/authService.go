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

type AuthService interface {
	UserRegisterService(user *model.User) (*model.UserView, error)
	UserEmailLoginService(credential *model.Credential) (*model.UserView, error)
	UserTokenLoginService(token *string) (*model.UserView, error)
}
