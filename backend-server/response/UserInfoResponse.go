package response

import "SimpleDouyin/module"

// UserInfo 获取用户信息返回体
type UserInfo struct {
	Response
	module.User `json:"user"`
}
