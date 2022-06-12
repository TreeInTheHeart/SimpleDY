package service

import (
	"SimpleDY/global"
	"SimpleDY/pojo"
	"SimpleDY/status"
	"time"
)

type VideoService struct {
}

// AddFavoriteCount
/*
增加点赞者数量
param :
	UserID  uint64 `json:"user_id"`
	VideoID uint64 `json:"video_id"`
	Token   string `json:"token"`
	Type    uint   `json:"action_type"`
*/

// GetNextId
/*为了确保用户上传的视频在本地存储时是唯一文件名，将文件名存储为 "id.mp3"格式
查询待上传视频id
*/
func (videoservice VideoService) GetNextId() uint64 {
	var count int64
	global.Db.Model(&pojo.Video{}).Count(&count)
	return uint64(count) + 1
}

// AddVideo
/*上传视频
param 视频路径 封面路径 标题 作者id
response 错误码
*/
func (videoservice VideoService) AddVideo(videoPath, coverPath, title string, authorId uint64) int64 {
	video := pojo.Video{
		AuthorId:      authorId,
		VideoPath:     videoPath,
		CoverPath:     coverPath,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		Status:        0,
		CreatedAt:     time.Now().Unix(),
	}
	if global.Db.Create(&video).RowsAffected == 1 {
		return 0
	}
	return status.UnknownError
}

//GetVideoListByAuthorId
/*通过作者id查询发布的视频列表
 */
func (videoservice VideoService) GetVideoListByAuthorId(authorId uint64) *pojo.VideoList {
	var videolist pojo.VideoList
	global.Db.Model(&pojo.Video{}).Where("authorId = ?", authorId).Find(&videolist)
	return &videolist
}

// FeedbyTime
// 返回截止时间Time时间之前的视频流 最多30条
// param 截止时间
// response 视频流
func (VideoService VideoService) FeedbyTime(time int64) *pojo.VideoList {
	var videolist pojo.VideoList
	global.Db.Model(&pojo.Video{}).Where("create_time < ?", time).Order("create_time DESC").Limit(3).Find(&videolist)
	return &videolist
}

//添加喜欢的数量
func (videoservice VideoService) AddFavoriteCount(param pojo.FavoritaParam) error {
	global.Db.Model(&pojo.Video{})
	return nil
}

//减少点赞者数量
func (videoservice VideoService) ReduceFavoriteCount() error {
	return nil
}

//修改点赞者数量 根据redis
func (v VideoService) UpdateFavoriteCount() error {
	return nil
}
