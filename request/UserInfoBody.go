package request

// UserInfoBody 获取用户信息请求体
type UserInfoBody struct {
	UserId int64 `json:"user_id"`
	Token  int64 `json:"token"`
}
