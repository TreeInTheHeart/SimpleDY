package service

import (
	"SimpleDY/global"
	"SimpleDY/pojo"
	"SimpleDY/status"
)

type UserService struct {
}

//Register
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

//Login
/**
param
response
*/
func (userservice UserService) Login(param pojo.UserLoginParam) (*pojo.User, uint64) {
	var user pojo.User
	var count int64
	global.Db.Model(&pojo.User{}).Where("username = ?", param.Username).Find(&user).Count(&count)
	if count == 0 {
		return nil, status.UserNotExistOrPasswordWrongError
	}

	if user.Password != param.Password {
		return nil, status.UserNotExistOrPasswordWrongError
	}
	return &user, status.Success
}

//GetInfoByUserId
func (userservice UserService) GetInfoByUserId(userid uint64) *pojo.User {
	var user pojo.User
	global.Db.Model(&pojo.User{}).Where("id = ?", userid).First(&user)
	return &user
}
