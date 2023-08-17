package auth

import "SimpleDouyin/response"

// UserRegister 注册结果返回体
type UserRegister struct {
	response.Response
	Token  string `json:"token"`   // 用户鉴权token
	UserId int64  `json:"user_id"` // 用户id
}

// NewUserRegister 生成一个成功的注册结果返回体
func NewUserRegister(userId int64, token string) UserRegister {
	return UserRegister{
		UserId: userId,
		Token:  token,
	}
}
