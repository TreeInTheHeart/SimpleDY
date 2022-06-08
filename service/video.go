package service

import (
	"SimpleDY/global"
	"SimpleDY/pojo"
	"SimpleDY/status"
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

/*查询上传视频id*/
func (videoservice VideoService) GetNextId() uint64 {
	var count int64
	global.Db.Model(&pojo.Video{}).Count(&count)
	return uint64(count) + 1
}

func (videoservice VideoService) AddVideo(videoPath, coverPath, title string, authorId uint64) int64 {
	video := pojo.Video{
		AuthorId:      authorId,
		VideoPath:     videoPath,
		CoverPath:     coverPath,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		Status:        0,
	}
	if global.Db.Create(&video).RowsAffected == 1 {
		return 0
	}
	return status.UnknownError
}

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
