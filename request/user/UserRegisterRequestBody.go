package user

// RegisterBody 用户注册请求体
type RegisterBody struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	CheckPassword string `json:"checkPassword"`
}
