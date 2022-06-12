package pojo

//Video表对应的数据库映射

type VideoList []Video
type Video struct {
	ID            uint64 `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	AuthorId      uint64 `gorm:"column:authorId"`
	VideoPath     string `json:"play_url" gorm:"column:video_path"`                     //视频路径
	CoverPath     string `json:"cover_url" gorm:"column:cover_path"`                    //封面地址
	FavoriteCount uint64 `json:"favorite_count" gorm:"column:favorite_count;default:0"` //点赞数
	CommentCount  uint64 `json:"comment_count" gorm:"column:comment_count;default:0"`   //评论数
	Title         string `json:"title" gorm:"column:title"`
	Status        uint64 `grom:"column:title"`                         //视频状态 分 审核中 正常 下架
	CreatedAt     int64  `json:"created_at" gorm:"column:create_time"` //投稿时间
}
