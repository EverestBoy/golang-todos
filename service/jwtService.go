/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package service

import (
	"todoGo/model"
)

/*
@author everestboy
*/

type JwtService interface {
	GetJwtToken(userId string) (*string, error)
	VerifyJwtToken(jwtToken string) (*model.JWTClaims, error)
}
