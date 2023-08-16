package service

import (
	"DouyinBackend/dao"
	"DouyinBackend/request"
	"DouyinBackend/response"
	"DouyinBackend/service/impl"
)

// NewUserService 工厂函数 构造service层实例 并传入所需的dao层实例
func NewUserService() UserService {
	return &impl.UserServiceImpl{
		UserDao: dao.NewUserDao(),
	}
}

// UserService 用户业务操作层
type UserService interface {

	// Register 注册
	Register(body request.UserRegisterBody) (response.UserRegister, error)

	// Login 登录
	Login(body request.UserLoginBody) (response.UserLogin, error)

	// Info 获取用户信息
	Info(body request.UserInfoBody) (response.UserInfo, error)
}
