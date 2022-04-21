/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"encoding/json"
	"fmt"
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
	todoService       = impl.NewTodoService()
	validationService = impl.NewValidationService()
	jwtServiceImpl    = impl.NewJwtService()
)

type handlerTodo struct{}

func NewTODOHandler() handler.TODOHandler {
	return &handlerTodo{}
}

func (c handlerTodo) GetTodo(resp http.ResponseWriter, req *http.Request) {
	// authenticating the user
	start := time.Now()
	log.Printf("Entered GetTodo")
	resp.Header().Set("Content-type", "application/json")
	reqToken := req.Header.Get("Authorization")
	if reqToken == "" {
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := model.ErrorTextMessage{"Token authorizatoin failed"}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	claim, err := jwtServiceImpl.VerifyJwtToken(reqToken)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := err.Error()
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	// getting user id from claim
	userId := claim.UserId

	// getting the id of todo
	id := req.URL.Query().Get("id")
	if id == "" {
		resp.WriteHeader(http.StatusNetworkAuthenticationRequired)
		errorMessage := model.ErrorTextMessage{ErrorMsg: "Missing id in parameter"}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}

	todo, error := todoService.FindTodo(&id, &userId)
	if error != nil {
		resp.WriteHeader(http.StatusNotFound)
		errorMessage := model.ErrorTextMessage{ErrorMsg: error.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	resp.WriteHeader(http.StatusOK)
	successResponse := model.TodoResponse{Message: "SUCCESS", Todo: todo}
	json.NewEncoder(resp).Encode(successResponse)
	log.Printf("Completed GetTodo , total time: %s", time.Now().Sub(start))
}

func (c handlerTodo) GetTodos(resp http.ResponseWriter, req *http.Request) {
	// authenticating the user
	start := time.Now()
	log.Printf("Entered GetTodos")
	resp.Header().Set("Content-type", "application/json")
	reqToken := req.Header.Get("Authorization")
	if reqToken == "" {
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := model.ErrorTextMessage{"Token authorizatoin failed"}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	claim, err := jwtServiceImpl.VerifyJwtToken(reqToken)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := err.Error()
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	// getting user id from claim
	userId := claim.UserId

	todos, error := todoService.FindAllTodos(&userId)
	if error != nil {
		resp.WriteHeader(http.StatusNetworkAuthenticationRequired)
		errorMessage := model.ErrorTextMessage{ErrorMsg: error.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	resp.WriteHeader(http.StatusOK)
	successResponse := model.TodosResponse{Message: "SUCCESS", Todos: todos}
	json.NewEncoder(resp).Encode(successResponse)
	log.Printf("Completed GetTodos , total time: %s", time.Now().Sub(start))
}

func (c handlerTodo) AddTodos(resp http.ResponseWriter, req *http.Request) {
	// authenticating the user
	start := time.Now()
	log.Printf("Entered AddTodos")
	resp.Header().Set("Content-type", "application/json")
	reqToken := req.Header.Get("Authorization")
	if reqToken == "" {
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := model.ErrorTextMessage{ErrorMsg: "token authorization failed"}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	claim, err := jwtServiceImpl.VerifyJwtToken(reqToken)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := err.Error()
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	// getting user id from claim
	userId := claim.UserId

	var todo model.TodoModel
	err = json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	log.Println("Validating todo")
	err = validationService.ValidateTodo(&todo)
	if err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	log.Println("Creating todo")
	createdTodo, err := todoService.CreateTodo(&todo, &userId)

	if err != nil {
		log.Printf("Got error %c", err)
		resp.WriteHeader(http.StatusNotFound)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	resp.WriteHeader(http.StatusCreated)
	successResponse := model.TodoResponse{Message: "SUCCESS", Todo: createdTodo}
	json.NewEncoder(resp).Encode(successResponse)
	log.Printf("Completed todoCreation , total time: %s", (time.Now().Sub(start)))
}

func (c handlerTodo) UpdateTodo(resp http.ResponseWriter, req *http.Request) {
	// authenticating the user
	start := time.Now()
	log.Printf("Entered UpdateTodo")
	resp.Header().Set("Content-type", "application/json")
	reqToken := req.Header.Get("Authorization")
	if reqToken == "" {
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := model.ErrorTextMessage{"Token authorizatoin failed"}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	claim, err := jwtServiceImpl.VerifyJwtToken(reqToken)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := err.Error()
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	// getting user id from claim
	userId := claim.UserId
	id := req.URL.Query().Get("id")
	if id == "" {
		resp.WriteHeader(http.StatusNetworkAuthenticationRequired)
		errorMessage := model.ErrorTextMessage{ErrorMsg: "Missing id in parameter"}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	var todo model.TodoModel
	err = json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	log.Println("Validating todo")
	err = validationService.ValidateTodo(&todo)
	if err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	log.Println("Updating todo")
	updatedTodo, err := todoService.UpdateTodo(&id, &userId, &todo)

	if err != nil {
		log.Printf("Got error %c", err)
		resp.WriteHeader(http.StatusNetworkAuthenticationRequired)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	resp.WriteHeader(http.StatusCreated)
	successResponse := model.TodoResponse{Message: "SUCCESS", Todo: updatedTodo}
	json.NewEncoder(resp).Encode(successResponse)
	log.Printf("Completed todoUpdate , total time: %s", (time.Now().Sub(start)))
}

func (c handlerTodo) DeleteTodo(resp http.ResponseWriter, req *http.Request) {
	// authenticating the user
	start := time.Now()
	log.Printf("Entered UpdateTodo")
	resp.Header().Set("Content-type", "application/json")
	reqToken := req.Header.Get("Authorization")
	if reqToken == "" {
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := model.ErrorTextMessage{"Token authorizatoin failed"}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	claim, err := jwtServiceImpl.VerifyJwtToken(reqToken)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := err.Error()
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	// getting user id from claim
	userId := claim.UserId

	id := req.URL.Query().Get("id")
	if id == "" {
		resp.WriteHeader(http.StatusNetworkAuthenticationRequired)
		errorMessage := model.ErrorTextMessage{ErrorMsg: "Missing id in parameter"}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}

	error := todoService.DeleteTodo(&id, &userId)
	if error != nil {
		resp.WriteHeader(http.StatusNotFound)
		errorMessage := model.ErrorTextMessage{ErrorMsg: error.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	resp.WriteHeader(http.StatusOK)
	var successResponse = model.MessageResponse{Message: fmt.Sprintf("Todo with id: %s is successfully deleted", id)}
	json.NewEncoder(resp).Encode(successResponse)
	log.Printf("Completed DeleteTodo , total time: %s", time.Now().Sub(start))
}
