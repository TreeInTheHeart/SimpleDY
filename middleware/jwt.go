package middleware

import (
	"SimpleDY/status"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var Salt string = "simpleDY"
var Issuer = "simpleDYGroup"

type MyClaims struct {
	UserId   uint64
	UserName string
	jwt.RegisteredClaims
}

func GenerateTokenString(userid uint64, usrname string) (string, error) {
	claims := MyClaims{
		UserId:   userid,
		UserName: usrname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			Issuer:    Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Salt))
}

func Parse(token string) (*MyClaims, bool) {

	tokenObj, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Salt), nil
	})
	if tokenObj == nil {
		return nil, false
	}
	if tokenObj.Valid {
		key, _ := tokenObj.Claims.(*MyClaims)
		return key, true
	} else {
		return nil, false
	}
}

func GetJwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取token
		tokenStr := c.Query("token")
		//未登录TokenIsNUll
		if tokenStr == "" {
			c.JSON(http.StatusOK, gin.H{"status_code": status.TokenIsNULL, "status_msg": status.Msg(status.TokenIsNULL)})
			c.Abort()
			return
		}

		//token解析错误
		tokenStruct, ok := Parse(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"status_code": status.TokenParseError, "status_msg": status.Msg(status.TokenParseError),
			})
			c.Abort()
			return
		}
		//超时
		if time.Now().Unix() > tokenStruct.ExpiresAt.Unix() {
			c.JSON(http.StatusOK, gin.H{
				"status_code": status.TokenIsExpired, "status_msg": status.Msg(status.TokenIsExpired),
			})
			c.Abort()
			return
		}
		c.Set("userid", tokenStruct.UserId)
		c.Set("user_name", tokenStruct.UserName)
		c.Next()

	}
}

func PostJwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, ok := c.GetPostForm("token")
		if !ok || tokenStr == "" {
			c.JSON(http.StatusOK, gin.H{
				"status_code": status.TokenIsNULL, "status_msg": status.Msg(status.TokenIsNULL),
			})
			c.Abort()
			return
		}
		//token解析错误
		tokenStruct, ok := Parse(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"status_code": status.TokenParseError, "status_msg": status.Msg(status.TokenParseError),
			})
			c.Abort()
			return
		}
		//超时
		if time.Now().Unix() > tokenStruct.ExpiresAt.Unix() {
			c.JSON(http.StatusOK, gin.H{
				"status_code": status.TokenIsExpired, "status_msg": status.Msg(status.TokenIsExpired),
			})
			c.Abort()
			return
		}
		c.Set("userid", tokenStruct.UserId)
		c.Set("user_name", tokenStruct.UserName)
		c.Next()
	}
}
