package user

// LoginBody 用户注册请求体
type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
