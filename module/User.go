package module

type User struct {
	Entity
	Username string `json:"username"`           // 账号
	Password string `json:"password"`           // 密码
	Nickname string `json:"nickname,omitempty"` // 昵称
	Gender   string `json:"gender,omitempty"`   // 性别
	Phone    string `json:"phone,omitempty"`    // 电话号码
	Email    string `json:"email,omitempty"`    // 邮箱
	Age      int    `json:"age,omitempty"`      // 年龄
	Status   int8   `json:"status,omitempty"`   // 状态
	Role     int8   `json:"role,omitempty"`     // 角色
}
