/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"todoGo/handler"
	"todoGo/model"
	"todoGo/service/impl"
)

/*
@author everetboy
*/

var (
	authService = impl.NewAuthService()
)

type handlerAuth struct{}

func NewAuthHandler() handler.AuthHandler {
	return &handlerAuth{}
}

func (h handlerAuth) UserEmailLogin(resp http.ResponseWriter, req *http.Request) {
	start := time.Now()
	log.Printf("Entered UserEmailLogin")
	resp.Header().Set("Content-type", "application/json")
	var credential model.Credential
	err := json.NewDecoder(req.Body).Decode(&credential)
	if err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	log.Println("Validating user")
	err = validationService.ValidateCredentials(&credential)
	if err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}

	log.Println("Validating user credentials")
	createdUser, err := authService.UserEmailLoginService(&credential)

	if err != nil {
		log.Printf("Got error %c", err)
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	resp.WriteHeader(http.StatusAccepted)
	successResponse := model.UserResponse{Message: "SUCCESS", User: createdUser}
	json.NewEncoder(resp).Encode(successResponse)
	log.Printf("Completed UserEmailLogin , total time: %s", (time.Now().Sub(start)))
}

func (h handlerAuth) UserTokenLogin(resp http.ResponseWriter, req *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h handlerAuth) UserRegister(resp http.ResponseWriter, req *http.Request) {
	start := time.Now()
	log.Printf("Entered UserRegister")
	resp.Header().Set("Content-type", "application/json")
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	log.Println("Validating user")
	err = validationService.ValidateUserRegister(&user)
	if err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	log.Println("Creating user")
	createdUser, err := authService.UserRegisterService(&user)

	if err != nil {
		log.Printf("Got error %c", err)
		resp.WriteHeader(http.StatusExpectationFailed)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	resp.WriteHeader(http.StatusCreated)
	successResponse := model.UserResponse{Message: "SUCCESS", User: createdUser}
	json.NewEncoder(resp).Encode(successResponse)
	log.Printf("Completed UserRegister , total time: %s", (time.Now().Sub(start)))
}
