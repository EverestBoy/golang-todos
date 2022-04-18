/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package repository

import "todoGo/model"

/*
@author everestboy
*/

type TODORepository interface {
	FindAll() ([]model.TodoModel, error)
}
