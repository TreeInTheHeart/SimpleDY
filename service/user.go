package service

import (
	"SimpleDY/global"
	"SimpleDY/model"
	"SimpleDY/utils"
	"SimpleDY/utils/json_response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type UserService struct {
}

func (userservice UserService) Register(c *gin.Context, user model.User) {
	if len(user.Name) > 32 || len(user.Password) > 32 {
		json_response.Error(c, -1, "用户名或密码过长")
		return
	}
	if len(user.Name) <= 0 || len(user.Password) < 5 {
		json_response.Error(c, -1, "用户名或密码过短")
		return
	}

	err := global.Db.Table("users").Where("name = ?", user.Name).Take(&user).Error
	if err == nil {
		json_response.Error(c, -1, "用户名已存在")
		return
	} else {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(err)
			json_response.Error(c, -1, "未知错误user.go_1")
			return
		}
	}

	passwordHashed, err := utils.HashPassword(user.Password)
	if err != nil {
		json_response.Error(c, -1, "register failed")
	}

	user.PasswordHashed = passwordHashed
	result := global.Db.Create(&user)
	if result.Error != nil {
		log.Println("register insert failed.", result.Error)
		json_response.Error(c, -1, "注册失败")
		return
	}

	token, err := utils.GenerateToken(user.Name, strconv.FormatInt(int64(user.ID), 10))
	if err != nil {
		json_response.Error(c, -1, "未知错误user.go_2")
		return
	}

	json_response.OK(c, "ok", model.UserLoginResponse{
		UserId: int64(user.ID),
		Token:  token,
	})

}

func (userservice UserService) Login(c *gin.Context, user model.User) {
	err := global.Db.Table("users").Where("name = ?", user.Name).Take(&user).Error

	if err != nil {
		json_response.Error(c, -1, "用户名不存在")
		return
	}

	// 使用密码的哈希值来验证
	if !utils.CheckPasswordHash(user.Password, user.PasswordHashed) {
		json_response.Error(c, -1, "登入失败")
		return
	}

	token, _ := utils.GenerateToken(user.Name, strconv.FormatInt(int64(user.ID), 10))

	json_response.OK(c, "ok", model.UserLoginResponse{
		UserId: int64(user.ID),
		Token:  token,
	})

}

func (userservice UserService) UseInfo(c *gin.Context, user model.User) {
	//err := global.Db.Table("users").Where("id = ?", c.MustGet("userID")).Take(&user).Error
	result := global.Db.Table("users").Where("name = ? AND password = ?", user.Name, user.Password).Take(&user)
	if result.Error == nil {
		json_response.Error(c, -1, "用户名不存在")
		return
	}
	followCount := user.FollowCount
	followerCount := user.FollowerCount
	json_response.OK(c, "ok", model.UserInfo{
		User:          user,
		FollowCount:   int64(followCount),
		FollowerCount: int64(followerCount),
	})
}
