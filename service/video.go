package service

import (
	"SimpleDY/global"
	"SimpleDY/pojo"
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
func (videoservice VideoService) AddFavoriteCount(param pojo.FavoritaParam) error {
	global.Db.Model(&pojo.Video{})
}

//减少点赞者数量
func (videoservice VideoService) ReduceFavoriteCount() error {

}

//修改点赞者数量 根据redis
func (v VideoService) UpdateFavoriteCount() error {
	return nil
}
