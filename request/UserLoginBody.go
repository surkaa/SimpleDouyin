package request

import "github.com/gin-gonic/gin"

// UserLoginBody 用户注册请求体
type UserLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ParseUserLoginBody 解析用户注册请求体
func ParseUserLoginBody(c *gin.Context) (UserLoginBody, error) {
	username := c.Query("username")
	password := c.Query("password")
	return UserLoginBody{
		Username: username,
		Password: password,
	}, nil
}
