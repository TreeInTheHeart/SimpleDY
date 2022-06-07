package handler

import (
	"SimpleDY/api"
	"SimpleDY/middleware"
	"github.com/gin-gonic/gin"
)

func Handler() {
	r := gin.Default()
	//authMiddleware := r.Group("/publish",middleware.JwtMiddleWare())
	r.POST("/user/register", api.Register)
	r.POST("/user/login", api.Login)
	r.GET("/user", middleware.JwtMiddleWare(), api.GetInfo)
	//r.POST("favorite/action", api.FavoriteOp)
	r.Run()
}
