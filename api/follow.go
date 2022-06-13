// Package api
// @author    : MuXiang123
// @time      : 2022/6/12 11:18
package api

import (
	"SimpleDY/middleware"
	"SimpleDY/pojo"
	"SimpleDY/service"
	"SimpleDY/status"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var followingService service.FollowingService

// RelationAction 关注/取消关注操作
func RelationAction(c *gin.Context) {
	//当前用户id
	strToken := c.Query("token")
	tokenStruct, _ := middleware.Parse(strToken)
	hostId := tokenStruct.UserId

	//获取待关注的用户id
	getToUserId, err2 := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	guestId := uint64(getToUserId)
	//获取关注操作（关注1，取消关注2）
	getActionType, err3 := strconv.ParseInt(c.Query("action_type"), 10, 64)
	actionType := uint(getActionType)
	if err2 != nil || err3 != nil {
		c.JSON(http.StatusOK, pojo.FollowResponse{
			StatusCode: status.UnknownError,
			StatusMsg:  status.Msg(status.UnknownError),
		})
	}
	//对自己关注/取消关注自己进行校验，不合法
	if guestId == hostId {
		c.JSON(http.StatusOK, pojo.FollowResponse{
			StatusCode: status.InabilityToFocusOnYourself,
			StatusMsg:  status.Msg(status.InabilityToFocusOnYourself),
		})
		c.Abort()
		return
	}

	//调用service层
	_, err := followingService.FollowAction(hostId, guestId, actionType)
	if err != nil {
		c.JSON(http.StatusBadRequest, pojo.FollowResponse{
			StatusCode: status.UnknownError,
			StatusMsg:  status.Msg(1),
		})
	} else {
		if actionType == 1 {
			c.JSON(http.StatusBadRequest, pojo.FollowResponse{
				StatusCode: status.Success,
				StatusMsg:  status.Msg(0),
			})
		}
		if actionType == 2 {
			c.JSON(http.StatusBadRequest, pojo.FollowResponse{
				StatusCode: status.Success,
				StatusMsg:  status.Msg(0),
			})
		}
	}
}

//FollowList 获取用户关注列表
func FollowList(c *gin.Context) {
	//获取用户本人id
	strToken := c.Query("token")
	tokenStruct, _ := middleware.Parse(strToken)
	hostId := tokenStruct.UserId
	fmt.Println(hostId, tokenStruct)
	//获取其他用户id
	getGuestId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	guestId := uint64(getGuestId)

	//判断查询类型，从数据库取用户列表
	var err error
	var userList []pojo.User
	if guestId == 0 {
		//若其他用户id为0，代表查本人的关注表
		userList, err = followingService.GetFollowingList(hostId)
	} else {
		//若其他用户id不为0，代表查对方的关注表
		userList, err = followingService.GetFollowingList(guestId)
	}

	//构造返回的数据
	var ReturnFollowerList = make([]pojo.ReturnFollower, len(userList))
	for i, m := range userList {
		ReturnFollowerList[i].Id = uint(m.Id)
		ReturnFollowerList[i].Name = m.Name
		ReturnFollowerList[i].FollowCount = uint(m.FollowCount)
		ReturnFollowerList[i].FollowerCount = uint(m.FollowerCount)
		ReturnFollowerList[i].IsFollow = service.IsFollowing(hostId, m.Id)
	}

	//响应返回, 粉丝表和关注表的返回结构体相同，复用
	if err != nil {
		c.JSON(http.StatusBadRequest, pojo.FollowingListResponse{
			StatusCode: 1,
			StatusMsg:  status.Msg(14),
			UserList:   nil,
		})
	} else {
		c.JSON(http.StatusOK, pojo.FollowingListResponse{
			StatusCode: 0,
			StatusMsg:  status.Msg(13),
			UserList:   ReturnFollowerList,
		})
	}
}

//FollowerList 获取用户粉丝列表
func FollowerList(c *gin.Context) {
	//获取用户本人id
	strToken := c.Query("token")
	tokenStruct, _ := middleware.Parse(strToken)
	hostId := tokenStruct.UserId
	//1.2获取其他用户id
	getGuestId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	guestId := uint(getGuestId)

	//判断查询类型
	var err error
	var userList []pojo.User
	if guestId == 0 {
		//查本人的粉丝表
		userList, err = followingService.GetFollowerList(hostId)
	} else {
		//查对方的粉丝表
		userList, err = followingService.GetFollowerList(uint64(guestId))
	}

	//3.判断查询类型，从数据库取用户列表
	var ReturnFollowerList = make([]pojo.ReturnFollower, len(userList))
	for i, m := range userList {
		ReturnFollowerList[i].Id = uint(m.Id)
		ReturnFollowerList[i].Name = m.Name
		ReturnFollowerList[i].FollowCount = uint(m.FollowCount)
		ReturnFollowerList[i].FollowerCount = uint(m.FollowerCount)
		ReturnFollowerList[i].IsFollow = service.IsFollowing(hostId, m.Id)
	}

	//响应返回, 粉丝表和关注表的返回结构体相同，复用
	if err != nil {
		c.JSON(http.StatusBadRequest, pojo.FollowingListResponse{
			StatusCode: status.UnknownError,
			StatusMsg:  status.Msg(14),
			UserList:   nil,
		})
	} else {
		c.JSON(http.StatusOK, pojo.FollowingListResponse{
			StatusCode: status.Success,
			StatusMsg:  status.Msg(13),
			UserList:   ReturnFollowerList,
		})
	}
}
