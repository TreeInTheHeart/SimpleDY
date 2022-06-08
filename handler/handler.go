package handler

import (
	"SimpleDY/api"
	"SimpleDY/middleware"
	"github.com/gin-gonic/gin"
)

func Handler() {
	r := gin.Default()
	r.MaxMultipartMemory = 10 << 20 //短视频大小最大不超过 10M
	//authMiddleware := r.Group("/publish",middleware.JwtMiddleWare())
	r.POST("/user/register", api.Register)
	r.POST("/user/login", api.Login)
	r.GET("/user", middleware.GetJwtMiddleWare(), api.GetInfo)
	//r.POST("favorite/action", api.FavoriteOp)
	r.POST("/publish/action", middleware.PostJwtMiddleWare(), api.Publish)
	r.Run()
}
