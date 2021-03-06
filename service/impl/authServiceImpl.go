/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"errors"
	"fmt"
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
	authRepo       = impl.NewAuthRepository()
	jwtServiceImpl = NewJwtService()
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
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
	// getting jwt token for user
	jwtToken, err := jwtServiceImpl.GetJwtToken(user.Id.Hex())
	if err != nil {
		return nil, err
	}
	// registering user
	createdUser, err := authRepo.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	// creating userview
	var userView = model.UserView{
		Username:  createdUser.Username,
		Phone:     createdUser.Phone,
		Address:   createdUser.Address,
		Email:     createdUser.Email,
		UserToken: model.Token{Token: *jwtToken},
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
	fmt.Println("Password")
	passwordMatch := CheckPasswordHash(credential.Password, user.Password)
	fmt.Println("Password")
	if !passwordMatch {
		return nil, failedAuth
	}

	// getting jwt token for user
	jwtToken, err := jwtServiceImpl.GetJwtToken(user.Id.Hex())
	if err != nil {
		return nil, err
	}

	// creating user view as user password matches
	var userView = model.UserView{
		Username:  user.Username,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		UserToken: model.Token{Token: *jwtToken},
	}
	return &userView, nil
}

func (a authService) UserTokenLoginService(token *string) (*model.UserView, error) {
	// verifying token
	claim, err := jwtServiceImpl.VerifyJwtToken(*token)
	if err != nil {
		return nil, err
	}
	userId := claim.UserId
	// getting user detail
	user, err := authRepo.UserDetailById(&userId)
	// getting jwt token for user
	jwtToken, err := jwtServiceImpl.GetJwtToken(user.Id.Hex())
	if err != nil {
		return nil, err
	}

	// creating user view as user password matches
	var userView = model.UserView{
		Username:  user.Username,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		UserToken: model.Token{Token: *jwtToken},
	}
	return &userView, nil

}
