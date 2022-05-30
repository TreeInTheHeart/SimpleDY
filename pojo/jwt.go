package pojo

import "github.com/golang-jwt/jwt/v4"

type MyClaims struct {
	Id       uint64
	UserName string
	jwt.RegisteredClaims
}
