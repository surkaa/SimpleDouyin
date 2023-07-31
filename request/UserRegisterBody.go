package request

// UserRegisterBody 用户注册请求体
type UserRegisterBody struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	CheckPassword string `json:"checkPassword"`
}
