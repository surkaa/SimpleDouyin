package response

// UserRegister 注册结果返回体
type UserRegister struct {
	Response
	UserId int64 `json:"user_id"`
	Token  int64 `json:"token"`
}
