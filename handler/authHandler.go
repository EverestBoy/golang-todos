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

type AuthHandler interface {
	UserEmailLogin(resp http.ResponseWriter, req *http.Request)
	UserTokenLogin(resp http.ResponseWriter, req *http.Request)
	UserRegister(resp http.ResponseWriter, req *http.Request)
}
