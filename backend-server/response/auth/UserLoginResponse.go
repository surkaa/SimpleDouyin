package auth

import "SimpleDouyin/response"

// UserLogin 注册结果返回体
type UserLogin struct {
	response.Response
	UserId int64 `json:"user_id"`
	Token  int64 `json:"token"`
}
