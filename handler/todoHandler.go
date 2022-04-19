/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package handler

import "net/http"

/*
@author everetboy
*/

type TODOHandler interface {
	GetTodos(resp http.ResponseWriter, req *http.Request)
	AddTodos(resp http.ResponseWriter, req *http.Request)
	GetTodo(resp http.ResponseWriter, req *http.Request)
	UpdateTodo(resp http.ResponseWriter, req *http.Request)
	DeleteTodo(resp http.ResponseWriter, req *http.Request)
}
