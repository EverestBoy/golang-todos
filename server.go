/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todoGo/handler/impl"
)

/*
@author everestboy
*/

var (
	todoHandler = impl.NewTODOHandler()
	authHandler = impl.NewAuthHandler()
)

func main() {
	router := mux.NewRouter()
	const port = ":8000"
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Up and runniing")
	})

	// auth handler
	router.HandleFunc("/login", authHandler.UserTokenLogin).Methods("GET")
	router.HandleFunc("/login", authHandler.UserEmailLogin).Methods("POST")
	router.HandleFunc("/register", authHandler.UserRegister).Methods("POST")

	// to-do handler
	router.HandleFunc("/todo", todoHandler.GetTodo).Methods("GET")
	router.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	router.HandleFunc("/todo", todoHandler.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo", todoHandler.AddTodos).Methods("POST")
	router.HandleFunc("/todo", todoHandler.DeleteTodo).Methods("DELETE")
	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
