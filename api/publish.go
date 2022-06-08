package api

import (
	"SimpleDY/pojo"
	"SimpleDY/service"
	"SimpleDY/status"
	"SimpleDY/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var videoService service.VideoService

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
