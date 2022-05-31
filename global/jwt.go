package global

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var Salt string = "simpleDY"
var Issuer = "simpleDYGroup"

type MyClaims struct {
	UserName string
	jwt.RegisteredClaims
}

func GenerateTokenString(usrname string) (string, error) {
	claims := MyClaims{
		UserName: usrname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			Issuer:    Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Salt))
}

func Parse(token string) (interface{}, error) {
	d1 := &MyClaims{}

	j, err := jwt.ParseWithClaims(token, d1, func(token *jwt.Token) (interface{}, error) {
		return []byte(Salt), nil
	})
	if err != nil {
		return
	}

	if j.Valid {
		return nil, nil
	}
}
