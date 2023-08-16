package request

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserInfoBody 获取用户信息请求体
type UserInfoBody struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

// ParseUserInfoBody 解析获取用户信息请求体
func ParseUserInfoBody(c *gin.Context) (UserInfoBody, error) {
	userId := c.Query("user_id")
	token := c.Query("token")
	userIdInt, _ := strconv.ParseInt(userId, 10, 64)
	return UserInfoBody{
		UserId: userIdInt,
		Token:  token,
	}, nil
}
