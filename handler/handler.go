package handler

import (
	"SimpleDY/api"
	"github.com/gin-gonic/gin"
)

func Handler() {
	r := gin.Default()
	r.POST("/user/register", api.Register)
	r.POST("/user/login", api.Login)
	//r.POST("favorite/action", api.FavoriteOp)
	r.Run()
}
