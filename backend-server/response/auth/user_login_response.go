package auth

import "SimpleDouyin/response"

// UserLogin 注册结果返回体
type UserLogin struct {
	response.Response
	Token  string `json:"token"`   // 用户鉴权token
	UserId int64  `json:"user_id"` // 用户id
}
