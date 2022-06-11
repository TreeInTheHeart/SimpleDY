package api

import (
	"SimpleDY/pojo"
	"SimpleDY/service"
	"SimpleDY/status"
	"SimpleDY/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var videoService service.VideoService

// Publish
/*发布新视频*/
func Publish(c *gin.Context) {
	file, err := c.FormFile("data")
	title, _ := c.GetPostForm("title")
	authorId, _ := c.Get("userid")

	if err != nil {
		c.JSON(http.StatusOK, pojo.PublishResponse{
			StatusCode: status.SaveUploadedFileError,
			StatusMsg:  status.Msg(status.SaveUploadedFileError),
		})
	}
	/*数据库中查询*/
	nextId := videoService.GetNextId()
	/*保存视频地址*/
	videoPath := utils.MakeVideoPathById(nextId, file.Filename)
	err = c.SaveUploadedFile(file, videoPath)
	if err != nil {
		c.JSON(http.StatusOK, pojo.PublishResponse{
			StatusCode: status.SaveUploadedFileError,
			StatusMsg:  status.Msg(status.SaveUploadedFileError),
		})
	}
	/*获得封面图片并保存地址*/
	coverPath := utils.MakeCoverPathById(nextId)
	err = utils.GetCoverFromVideo(videoPath, coverPath)
	if err != nil {
		c.JSON(http.StatusOK, pojo.PublishResponse{
			StatusCode: status.SaveUploadedFileError,
			StatusMsg:  status.Msg(status.SaveUploadedFileError),
		})
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"coverPath": coverPath,
	//	"videoPath": videoPath,
	//})
	/*存储到数据库*/
	ErroeCode := videoService.AddVideo(videoPath, coverPath, title, authorId.(uint64))
	//
	c.JSON(http.StatusOK, pojo.PublishResponse{
		StatusCode: ErroeCode,
		StatusMsg:  status.Msg(int(ErroeCode)),
	})

	//c.JSON(http.StatusOK, pojo.PublishResponse{
	//	StatusCode: 0,
	//	StatusMsg:  nil,
	//})
}

// GetPublishListByAuthorId
/*根据用户ID查询发布视频列表*/
func GetPublishListByAuthorId(c *gin.Context) {
	userIDstr := c.Query("user_id")
	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, pojo.GetVideoListResponse{
			StatusCode: status.UnknownError,
			StatusMsg:  status.Msg(status.UnknownError),
			VideoList:  nil,
		})
	}
	videoList := videoService.GetVideoListByAuthorId(userID)
	userInfo := userService.GetInfoByUserId(userID)

	videoRespnseList := make([]pojo.VideoResponse, len(*videoList))

	for idx, elem := range *videoList {
		//还应查询点赞关系表 填充IsFavorite这项
		videoRespnseList[idx] = pojo.VideoResponse{
			Author: pojo.Author{
				FollowCount:   int64(userInfo.FollowCount),
				FollowerCount: int64(userInfo.FollowerCount),
				ID:            int64(userInfo.Id),
				IsFollow:      userInfo.IsFollow,
				Name:          userInfo.Name,
			},
			CommentCount:  int64(elem.CommentCount),
			CoverURL:      elem.CoverPath,
			FavoriteCount: int64(elem.FavoriteCount),
			ID:            int64(elem.ID),
			IsFavorite:    false,
			PlayURL:       elem.VideoPath,
			Title:         elem.Title,
		}
	}
	c.JSON(http.StatusOK, pojo.GetVideoListResponse{
		StatusCode: 0,
		StatusMsg:  status.Msg(0),
		VideoList:  videoRespnseList,
	})
}
