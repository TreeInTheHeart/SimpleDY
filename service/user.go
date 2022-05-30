package service

import (
	"SimpleDY/global"
	"SimpleDY/pojo"
	"SimpleDY/status"
)

type UserService struct {
}

/**
param:用户注册信息包括用户名，密码，昵称
response:注册结果，用户id,错误码
*/
func (userservice UserService) Register(param pojo.UserRegisterParam) (bool, uint64, int) {
	var count int64
	global.Db.Model(&pojo.User{}).Where("username = ?", param.Username).Count(&count)
	if count > 0 {
		return false, 0, status.UsernameHasExistedError
	}
	user := pojo.User{
		Name:     param.Name,
		Username: param.Username,
		Password: param.Password,
	}
	if global.Db.Create(&user).RowsAffected == 1 {
		return true, user.Id, 0
	}
	return false, 0, status.UnknownError
}

/**
param
response
*/
func (userservice UserService) Login(param pojo.UserLoginParam) (bool, uint64) {
	var user pojo.User
	var count int64
	global.Db.Model(&pojo.User{}).Where("username = ?", param.Username, param.Password).Find(&user).Count(&count)
	if count == 0 {
		return
	}

	if user.Password != param.Password {
		return false, status.UserNotExistOrPasswordWrongError
	}
}
