package pojo

//Video表对应的数据库映射
type Video struct {
	ID            uint64 `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	AuthorId      uint64 `gorm:"column:authorId"`
	VideoPath     string `json:"play_url" gorm:"column:video_path"`                     //视频路径
	CoverPath     string `json:"cover_url" gorm:"column:cover_path"`                    //封面地址
	FavoriteCount uint64 `json:"favorite_count" gorm:"column:favorite_count;default:0"` //点赞数
	CommentCount  uint64 `json:"comment_count" gorm:"column:comment_count;default:0"`   //评论数
	Title         string `json:"title" gorm:"column:title"`
	Status        uint64 `grom:"column:title"` //视频状态 分 审核中 正常 下架
}

type author struct {
	ID            uint   `json:"id" `
	Name          string `json:"name"`
	FollowCount   uint64 `json:"follow_count"`
	FollowerCount uint64 `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

//返回视频列表信息
type VideoResponse struct {
	Video
	Author author
}
