package request

import "github.com/gin-gonic/gin"

// UserRegisterBody 用户注册请求体
type UserRegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// CheckPassword string `json:"checkPassword"` TODO 未来可试
}

// ParseUserRegisterBody 解析请求体中的参数得到注册请求体
func ParseUserRegisterBody(c *gin.Context) (UserRegisterBody, error) {
	username := c.Query("username")
	password := c.Query("password")
	return UserRegisterBody{
		Username: username,
		Password: password,
	}, nil
}
