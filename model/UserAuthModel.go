/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*
@author everestboy
*/

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Phone     string             `json:"phone" bson:"phone"`
	Address   string             `json:"address" bson:"address"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type Token struct {
	Token string `json:"token"`
}

type UserView struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	UserToken Token  `json:"jwtToken"`
}

type Credential struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
