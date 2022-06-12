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
	r.POST("/douyin/user/register/", api.Register)
	r.POST("/douyin/user/login/", api.Login)
	r.GET("/douyin/user/", middleware.GetJwtMiddleWare(), api.GetInfo)
	//r.POST("favorite/action", api.FavoriteOp)
	r.POST("/douyin/publish/action/", middleware.PostJwtMiddleWare(), api.Publish)
	r.GET("/douyin/publish/list/", middleware.GetJwtMiddleWare(), api.GetPublishListByAuthorId)
	r.GET("/douyin/feed/", api.Feed)
	r.Run()
}
