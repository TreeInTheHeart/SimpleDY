package initial

import (
	"SimpleDY/global"
	"SimpleDY/pojo"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Mysql() {
	m := global.Config.Mysql
	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gorm open err: %v\n", err)
		os.Exit(1)
	}
	if err := db.AutoMigrate(&pojo.User{}); err != nil {
		fmt.Fprintf(os.Stderr, "AutoMigrate err: %v\n", err)
		os.Exit(1)
	}
	global.Db = db
}
