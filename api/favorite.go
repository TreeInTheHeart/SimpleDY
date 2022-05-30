package api

import (
	"SimpleDY/pojo"
	"SimpleDY/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

//点赞操作
func FavoriteOp(c *gin.Context) {
	var param pojo.FavoritaParam
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusOK, pojo.Response{
			Code: status.RequestParamError,
			Msg:  status.Msg(status.RequestParamError),
		})
		return
	}
	//更新user_like_video表

	//增减video表中点赞者数量

}
