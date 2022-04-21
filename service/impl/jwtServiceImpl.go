/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
	"todoGo/model"
	"todoGo/service"
)

type jwtService struct{}

var (
	claims = model.JWTClaims{}
)

func NewJwtService() service.JwtService {
	return &jwtService{}
}

// Create the JWT key used to create the signature
var jwtKey = []byte("sdfjkj32432kl4j")

func (j jwtService) GetJwtToken(userId string) (*string, error) {

	claims.UserId = userId
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Minute * 30).Unix()

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	tokenString = "Bearer " + tokenString
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return nil, err
	}

	return &tokenString, nil
}

func (j jwtService) VerifyJwtToken(jwtToken string) (*model.JWTClaims, error) {
	realToken := strings.Split(jwtToken, "Bearer ")[1]
	println(realToken)
	tkn, err := jwt.ParseWithClaims(realToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		println(err.Error())
		if err == jwt.ErrSignatureInvalid {
			//w.WriteHeader(http.StatusUnauthorized)
			return nil, errors.New("not authorized")
		} else {
			return nil, errors.New("bad request")
		}
		//w.WriteHeader(http.StatusBadRequest)
		//return
	}
	if !tkn.Valid {
		//w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("not authorized")
	}

	payload, ok := tkn.Claims.(*model.JWTClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return payload, nil
}
