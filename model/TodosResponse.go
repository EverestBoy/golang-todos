/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package model

/*
@author everestboy
*/
type TodosResponse struct {
	Message string      `json:"message"`
	Todos   []TodoModel `json:"todos"`
}

type TodoResponse struct {
	Message string     `json:"message"`
	Todo    *TodoModel `json:"todo"`
}
type UserResponse struct {
	Message string    `json:"message"`
	User    *UserView `json:"user"`
}
