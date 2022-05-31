package api

import (
	"SimpleDY/global"
	"SimpleDY/pojo"
	"SimpleDY/service"
	"SimpleDY/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userService service.UserService

func Register(c *gin.Context) {
	var param pojo.UserRegisterParam
	err := c.ShouldBind(&param)

	/* 解析参赛错误 */
	if err != nil {
		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
			Token:      "",
			UserID:     0,
		})
	}

	//生成token
	token, err := global.GenerateTokenString(param.Username)
	if err != nil {
		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: status.GenerateTokenError,
			StatusMsg:  status.Msg(status.GenerateTokenError),
			Token:      "",
			UserID:     0,
		})
	}

	//注册
	ok, userId, errCode := userService.Register(param)

	if ok { //注册成功
		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: int64(errCode),
			StatusMsg:  status.Msg(0),
			Token:      token,
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
	err := c.ShouldBind(&param)
	/* 解析参赛错误 */
	if err != nil {
		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
			Token:      "",
			UserID:     0,
		})
	}
	user, code := userService.Login(param)
	if user != nil {
		//生成token
		token, err := global.GenerateTokenString(param.Username)
		if err != nil {
			c.JSON(http.StatusOK, pojo.UserLoginResponse{
				StatusCode: status.GenerateTokenError,
				StatusMsg:  status.Msg(status.GenerateTokenError),
				Token:      "",
				UserID:     0,
			})
		} else {
			c.JSON(http.StatusOK, pojo.UserLoginResponse{
				StatusCode: int64(code),
				StatusMsg:  status.Msg(0),
				Token:      token,
				UserID:     int64(user.Id),
			})
		}
	} else {
		c.JSON(http.StatusOK, pojo.UserLoginResponse{
			StatusCode: int64(code),
			StatusMsg:  status.Msg(int(code)),
			Token:      "",
			UserID:     0,
		})
	}
}
