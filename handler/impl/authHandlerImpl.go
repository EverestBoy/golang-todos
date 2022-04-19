/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"net/http"
	"todoGo/handler"
)

/*
@author everetboy
*/

type handlerAuth struct{}

func NewAuthHandler() handler.AuthHandler {
	return &handlerAuth{}
}

func (h handlerAuth) UserEmailLogin(resp http.ResponseWriter, req *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h handlerAuth) UserTokenLogin(resp http.ResponseWriter, req *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h handlerAuth) UserRegister(resp http.ResponseWriter, req *http.Request) {
	//TODO implement me
	panic("implement me")
}
