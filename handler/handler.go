package handler

import (
	"SimpleDY/api"
	"github.com/gin-gonic/gin"
)

func Handler()  {
	r :=gin.Default()
	r.POST("/user/register",api.Register)
	r.Run()
}