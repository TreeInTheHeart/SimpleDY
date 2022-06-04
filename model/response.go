package model

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfo struct {
	User          `json:"user"`
	FollowCount   int64 `json:"follow_count,omitempty"`
	FollowerCount int64 `json:"follower_count,omitempty"`
	IsFollow      bool  `json:"is_follow,omitempty"`
}
