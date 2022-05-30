package pojo

//点赞post操作获取参数
type FavoritaParam struct {
	UserID  uint64 `json:"user_id"`
	VideoID uint64 `json:"video_id"`
	Token   string `json:"token"`
	Type    uint   `json:"action_type"`
}

//user_like_video表
//存储用户点赞视频的相关信息 提升性能可以redis实现
type UserLikeVideo struct {
}
