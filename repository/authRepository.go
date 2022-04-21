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

type AuthRepository interface {
	RegisterUser(user *model.User) (*model.User, error)
	UserDetail(email *string, username *string) (*model.User, error)
	UserDetailById(userId *string) (*model.User, error)
}
