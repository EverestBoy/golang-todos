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
	"todoGo/controller"
)

/*
@author everestboy
*/

var (
	todoController controller.TODOController = controller.NewTODOController()
)

func main() {
	router := mux.NewRouter()
	const port = ":8000"
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Up and runniing")
	})

	router.HandleFunc("/todos", todoController.GetTodos).Methods("GET")
	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
