package impl

import (
	"DouyinBackend/dao"
	"DouyinBackend/db"
	"DouyinBackend/module"
	"DouyinBackend/request"
	"DouyinBackend/response"
)

// UserServiceImpl 用户业务层实现类
type UserServiceImpl struct {
	UserDao dao.UserDao
}

// Register 注册
func (userServiceImpl UserServiceImpl) Register(body request.UserRegisterBody) (response.UserRegister, error) {
	// 先检查是否已有这个账号:
	user, err := userServiceImpl.UserDao.GetUserByUsername(body.Username)
	if err != nil {
		return response.UserRegister{}, err
	}
	user = module.User{
		Username: body.Username,
		Password: body.Password,
	}
	mysqlDB := db.GetMysqlConnection()
	mysqlDB.Create(&user)
	return response.UserRegister{}, nil
}

// Login 登录
func (userServiceImpl UserServiceImpl) Login(body request.UserLoginBody) (response.UserLogin, error) {
	return response.UserLogin{}, nil
}

// Info 获取用户信息
func (userServiceImpl UserServiceImpl) Info(body request.UserInfoBody) (response.UserInfo, error) {
	return response.UserInfo{}, nil
}
