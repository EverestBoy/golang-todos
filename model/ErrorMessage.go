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
type ErrorMessage struct {
	ErrorMsg error `json:"errorMsg"`
}
type ErrorTextMessage struct {
	ErrorMsg string `json:"error_msg"`
}
