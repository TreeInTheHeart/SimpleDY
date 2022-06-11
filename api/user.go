package api

import (
	"SimpleDY/middleware"
	"SimpleDY/pojo"
	"SimpleDY/service"
	"SimpleDY/status"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userService service.UserService

func Register(c *gin.Context) {
	var param pojo.UserRegisterParam
	err := c.ShouldBind(&param)

	/* 解析参赛错误 */
	if err != nil {
		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
			Token:      "",
			UserID:     0,
		})
	}

	//密码加盐

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(param.Password))
	pw := hex.EncodeToString(md5Ctx.Sum(nil))

	//注册
	ok, userId, errCode := userService.Register(param.Username, pw)

	if ok { //注册成功

		//生成token
		token, err := middleware.GenerateTokenString(userId, param.Username)
		if err != nil {
			c.JSON(http.StatusOK, pojo.UserRegisterResponse{
				StatusCode: status.GenerateTokenError,
				StatusMsg:  status.Msg(status.GenerateTokenError),
				Token:      "",
				UserID:     0,
			})
			//token生成失败 是否需要在数据库中把刚刚注册的用户删除

		}

		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: int64(errCode),
			StatusMsg:  status.Msg(0),
			Token:      token,
			UserID:     userId,
		})
	} else {
		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: int64(errCode),
			StatusMsg:  status.Msg(errCode),
			Token:      "",
			UserID:     userId,
		})
	}
}

func Login(c *gin.Context) {
	var param pojo.UserLoginParam
	err := c.ShouldBind(&param)
	/* 解析参赛错误 */
	if err != nil {
		c.JSON(http.StatusOK, pojo.UserRegisterResponse{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
			Token:      "",
			UserID:     0,
		})
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(param.Password))
	pw := hex.EncodeToString(md5Ctx.Sum(nil))

	fmt.Println(pw)
	userID, code := userService.Login(param.Username, pw)
	if userID != 0 {
		//生成token
		token, err := middleware.GenerateTokenString(userID, param.Username)
		if err != nil {
			c.JSON(http.StatusOK, pojo.UserLoginResponse{
				StatusCode: status.GenerateTokenError,
				StatusMsg:  status.Msg(status.GenerateTokenError),
				Token:      "",
				UserID:     0,
			})
		} else {
			c.JSON(http.StatusOK, pojo.UserLoginResponse{
				StatusCode: int64(code),
				StatusMsg:  status.Msg(0),
				Token:      token,
				UserID:     int64(userID),
			})
		}
	} else {
		c.JSON(http.StatusOK, pojo.UserLoginResponse{
			StatusCode: int64(code),
			StatusMsg:  status.Msg(int(code)),
			Token:      "",
			UserID:     0,
		})
	}
}

func GetInfo(c *gin.Context) {
	userid, ok := c.Get("userid")
	//参数解析错误
	if !ok {
		c.JSON(http.StatusOK, pojo.GetUserInfoResponse{
			StatusCode: status.RequestParamError,
			StatusMsg:  status.Msg(status.RequestParamError),
			UserInfo:   pojo.UserInfo{},
		})
	}
	user := userService.GetInfoByUserId(userid.(uint64))
	c.JSON(http.StatusOK, pojo.GetUserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "Success",
		UserInfo: pojo.UserInfo{
			FollowCount:   int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
			ID:            int64(user.Id),
			IsFollow:      user.IsFollow,
			Name:          user.Name,
		},
	})
}
