package auth

import (
	"SimpleDouyin/response"
	"SimpleDouyin/response/data-module"
)

// UserInfo 获取用户信息返回体
type UserInfo struct {
	response.Response
	User *data_module.User `json:"user"` // 用户信息
}
