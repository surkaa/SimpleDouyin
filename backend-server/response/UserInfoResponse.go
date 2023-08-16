package response

import "DouyinBackend/module"

// UserInfo 获取用户信息返回体
type UserInfo struct {
	Response
	module.User `json:"user"`
}
