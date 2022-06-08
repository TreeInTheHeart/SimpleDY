package global

import (
	"SimpleDY/config"
	"gorm.io/gorm"
)

var (
	Config         config.Config
	Db             *gorm.DB
	VideoLibPath   string = "./data/video/"
	VideoCoverPath string = "./data/cover/"
)
