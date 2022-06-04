package model

////User数据库映射
//type User struct {
//	Id            uint64 `gorm:"column:id;autoIncrement;primaryKey"`
//	Name          string `gorm:"column:name"`
//	Username      string `gorm:"column:username"`
//	Password      string `gorm:"column:password"`
//	FollowCount   uint64 `gorm:"column:follow_count"`
//	FollowerCount uint64 `gorm:"column:follower_count"`
//	IsFollow      bool   `gorm:"column:is_follow"`
//}
//
////用户注册参数
//type UserRegisterParam struct {
//	Username string `form:"username" json:"username"`
//	Password string `form:"password" json:"password"`
//	Name string  `form:"name" json:"name"`
//}
//
////注册返回信息
//type UserRegisterResponse struct {
//	Response
//	UserId uint64 `json:"id"`
//}
//
////用户登录参数
//type UserLoginParam struct {
//	Username string `form:"username" json:"username"`
//	Password string `form:"password" json:"password"`
//}
