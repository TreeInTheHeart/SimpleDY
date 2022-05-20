package api

import (
	"SimpleDY/pojo"
	"SimpleDY/service"
	"SimpleDY/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userService service.UserService
func Register(c *gin.Context)  {
	var param pojo.UserRegisterParam
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusOK,pojo.UserRegisterResponse{
			Response: pojo.Response{
				Code: status.RequestParamError,
				Msg: status.Msg(status.RequestParamError),
			},
		})
		return
	}
	//注册
	ok,userId := userService.Register(param)
	if ok{
		c.JSON(http.StatusOK,userId)
	}else{
		c.JSON(http.StatusOK,"注册失败")
	}
}
