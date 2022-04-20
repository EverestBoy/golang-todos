/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"todoGo/model"
	"todoGo/repository/impl"
	"todoGo/service"
)

type authService struct{}

func NewAuthService() service.AuthService {
	return &authService{}
}

var (
	authRepo = impl.NewAuthRepository()
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a authService) UserRegisterService(user *model.User) (*model.UserView, error) {
	// hashing the password
	password, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password

	// creating user
	currentTime := time.Now()
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime
	user.Id = primitive.NewObjectID()
	createdUser, err := authRepo.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	// creating userview
	var userView = model.UserView{
		Username: createdUser.Username,
		Phone:    createdUser.Phone,
		Address:  createdUser.Address,
		Email:    createdUser.Email,
	}
	return &userView, nil

}

func (a authService) UserEmailLoginService(credential *model.Credential) (*model.UserView, error) {
	failedAuth := errors.New("user authentication failed")

	// getting user detail
	var email = credential.Email
	var username = credential.Username
	user, err := authRepo.UserDetail(&email, &username)
	log.Println(err)
	if err != nil {
		return nil, failedAuth
	}

	// checking password as user exists
	passwordMatch := CheckPasswordHash(credential.Password, user.Password)
	if !passwordMatch {
		return nil, failedAuth
	}

	// creating user view as user password matches
	var userView = model.UserView{
		Username: user.Username,
		Email:    user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
	}
	return &userView, nil
}

func (a authService) UserTokenLoginService(token *string) (*model.UserView, error) {
	//TODO implement me
	panic("implement me")
}
