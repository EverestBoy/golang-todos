/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package controller

import (
	"encoding/json"
	"net/http"
	"todoGo/model"
	service "todoGo/service"
)

type ErrorMessage struct {
	errorMsg string
}

/*
@author everetboy
*/

var (
	todoService service.TodoService = service.NewTodoService()
)

type TODOController interface {
	GetTodos(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewTODOController() TODOController {
	return &controller{}
}

func (c controller) GetTodos(resp http.ResponseWriter, req *http.Request) {
	print("Getting todos")
	resp.Header().Set("Content-type", "application/json")
	todos, error := todoService.FindAllTodos()
	if error != nil {
		resp.WriteHeader(http.StatusNetworkAuthenticationRequired)
		errorMessage := model.ErrorMessage{ErrorMsg: error}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	resp.WriteHeader(http.StatusOK)
	successResponse := model.TodosResponse{Message: "SUCCESS", Todos: todos}
	json.NewEncoder(resp).Encode(successResponse)
}
