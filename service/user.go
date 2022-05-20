package service

import (
	"SimpleDY/global"
	"SimpleDY/pojo"
)

type UserService struct {
}
/**
	param:用户注册信息包括用户名，密码，昵称
	response:注册结果，用户id
 */
func (userservice UserService) Register(param pojo.UserRegisterParam) (bool, uint64) {
	var count int64
	global.Db.Model(&pojo.User{}).Where("username = ?", param.Username).Count(&count)
	if count > 0 {
		return false, 0
	}
	user := pojo.User{
		Name:     param.Name,
		Username: param.Username,
		Password: param.Password,
	}
	if global.Db.Create(&user).RowsAffected == 1 {
		return true, user.Id
	}
	return false, 0
}
