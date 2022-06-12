package api

import (
	"SimpleDY/pojo"
	"SimpleDY/status"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {
	//时间返回的视频流的发布时间需要早于time
	timeStr := c.Query("latest_time")

	//token := c.Query("token")
	t, _ := strconv.ParseInt(timeStr, 10, 64)
	if t == 0 {
		t = time.Now().Unix()
	}
	//从数据库查到的视频列表
	videoList := videoService.FeedbyTime(t)
	//处理为视频返回信息列表
	VideoResponseList := make([]pojo.VideoResponse, len(*videoList))

	for idx, elem := range *videoList {
		author := userService.GetInfoByUserId(elem.AuthorId)
		VideoResponseList[idx] = pojo.VideoResponse{
			Author: pojo.Author{
				FollowCount:   int64(author.FollowCount),
				FollowerCount: int64(author.FollowerCount),
				ID:            int64(author.Id),
				IsFollow:      false,
				Name:          author.Name,
			},
			CommentCount:  int64(elem.CommentCount),
			CoverURL:      elem.CoverPath,
			FavoriteCount: int64(elem.FavoriteCount),
			ID:            int64(elem.ID),
			//IsFavorite状态和当前登录用户有关  此处应有另一张关系表 和每个用户对应  或者用redis实现这种关系的存储 ...总之这里没有实现 希望之后能完善
			IsFavorite: false,
			PlayURL:    elem.VideoPath,
			Title:      elem.Title,
		}
	}

	lastTime := (*videoList)[0].CreatedAt

	c.JSON(http.StatusOK, pojo.FeedResponse{
		NextTime:   lastTime,
		StatusCode: 0,
		StatusMsg:  status.Msg(0),
		VideoList:  VideoResponseList,
	})
}
