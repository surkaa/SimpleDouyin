package request

// UserLoginBody 用户注册请求体
type UserLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
