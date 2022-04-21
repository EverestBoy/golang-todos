/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package model

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
