package pojo

//点赞post操作获取参数
type FavoritaParam struct {
	UserID  uint64 `json:"user_id"`
	VideoID uint64 `json:"video_id"`
	Token   string `json:"token"`
	Type    uint   `json:"action_type"`
}

//user_like_video表
//存储用户点赞视频的纪律 提升性能可以redis实现
type UserLikeVideo struct {
	UserID  uint64 `gorm:"user_id"`
	VideoID int64  `gorm:"video_id"`
}
