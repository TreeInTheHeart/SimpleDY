// Package pojo
// @author    : MuXiang123
// @time      : 2022/6/11 16:31
package pojo

// Follow 关注数据库映射
// HostId 关注者
// GuestId 被关注者
type Follow struct {
	Id      uint64 `gorm:"column:id;autoIncrement;primaryKey"`
	HostId  uint64 `gorm:"column:host_id"`
	GuestId uint64 `gorm:"column:guest_id"`
}
