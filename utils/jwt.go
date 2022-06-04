package utils

import (
	"SimpleDY/global"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(username string, userid string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: username,
		UserID:   userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	tokenString, err := claims.SignedString(global.TokenSecret)

	return tokenString, err
}

func CheckToken(tokenString string) (*Claims, error) {
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return global.TokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// TODO: 动态配置时间
	if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) <= 0 {
		return claim, errors.New("the token has expired,please login again")
	}

	return claim, nil
}
