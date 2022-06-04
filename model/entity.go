package model

import (
	"gorm.io/gorm"
	"time"
)

// TODO: https://gorm.io/zh_CN/docs/indexes.html

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Video struct {
	Model
	// TODO: HashID
	AuthorID int64  `json:"author_id" `
	Author   User   `json:"author" `
	PlayUrl  string `json:"play_url"            `
	CoverUrl string `json:"cover_url,omitempty"   `
}

type Comment struct {
	Model

	VideoId   int64 `json:"video_id,omitempty"`
	UserID    int64
	User      User      `json:"user" `
	Content   string    `json:"content,omitempty" `
	CreatedAt time.Time `json:"create_date"`
}

type User struct {
	Model
	ID   uint   `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name string `json:"name"`
	// TODO: ignore Password in JSON
	Password       string `json:"-" `
	PasswordHashed string `json:"-" `
	FollowCount    uint   `json:"follow_count"`
	FollowerCount  uint   `json:"follower_count"`
	IsFollow       bool   `gorm:"-" json:"is_follow"`
}

type Follow struct {
	Model

	Name       string `json:"name,omitempty"`
	FollowerId int64  `json:"follower_id,omitempty" ` // 关注人
	FolloweeId int64  `json:"followee_id,omitempty" ` // 被关注人
	IsFollow   bool   `json:"is_follow,omitempty"`
}
