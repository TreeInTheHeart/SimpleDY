package api

import (
	"SimpleDY/pojo"
	"SimpleDY/service"
	"SimpleDY/status"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/gin-gonic/gin"
	"net/http"
)

var userService service.UserService

func Register(c *gin.Context) {
	var param pojo.UserRegisterParam
	err := c.ShouldBind(&param)

	//if err != nil {
	//	c.JSON(http.StatusOK, pojo.UserRegisterResponse{
	//		Response: pojo.Response{
	//			Code: status.RequestParamError,
	//			Msg:  status.Msg(status.RequestParamError),
	//		},
	//	})
	//	return
	//}

	if err != nil {
		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
			Token:      "",
			UserID:     -1,
		})
	}

	//注册
	ok, userId, errCode := userService.Register(param)

	if ok {
		//获取token
		claims := pojo.MyClaims{
			Id:       userId,
			UserName: param.Name,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
				Issuer:    "simpleDY",
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: int64(errCode),
			StatusMsg:  status.Msg(0),
			Token:      "",
			UserID:     userId,
		})
	} else {
		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: int64(errCode),
			StatusMsg:  status.Msg(errCode),
			Token:      "",
			UserID:     userId,
		})
	}
}

func Login(c *gin.Context) {
	var param pojo.UserLoginParam

}
