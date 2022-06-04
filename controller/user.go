package controller

import (
	"SimpleDY/model"
	"SimpleDY/service"
	"github.com/gin-gonic/gin"
)

var userService service.UserService

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user model.User // model\entity.go\User
	user = model.User{
		Name:     username,
		Password: password,
	}

	userService.Register(c, user) // service\user.go\Register
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user model.User // model\entity.go\User
	user = model.User{
		Name:     username,
		Password: password,
	}
	userService.Login(c, user) // service\user.go\Login
}

func UserInfo(c *gin.Context) {
	var user model.User          // model\entity.go\User
	userService.UseInfo(c, user) // service\user.go\UseInfo
}
