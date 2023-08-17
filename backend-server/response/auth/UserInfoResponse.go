package auth

import (
	"SimpleDouyin/module"
	"SimpleDouyin/response"
)

// UserInfo 获取用户信息返回体
type UserInfo struct {
	response.Response
	module.User `json:"user"`
}

func test() {
	var info = UserInfo{}
	unmarshalResponse, err := response.UnmarshalResponse([]byte{})
	if err != nil {
		return
	}
}
