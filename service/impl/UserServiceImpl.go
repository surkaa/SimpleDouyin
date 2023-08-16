package impl

import (
	"DouyinBackend/dao"
	"DouyinBackend/db"
	"DouyinBackend/module"
	"DouyinBackend/request"
	"DouyinBackend/response"
	"errors"
	"gorm.io/gorm"
	"regexp"
)

// UserServiceImpl 用户业务层实现类
type UserServiceImpl struct {
	UserDao dao.UserDao
}

const (
	// UsernameMinLen 账号最小长度
	UsernameMinLen = 1

	// UsernameMaxLen 账号最长长度
	UsernameMaxLen = 10

	// PasswordMinLen 密码最小长度
	PasswordMinLen = 6

	// PasswordMaxLen 密码最长长度
	PasswordMaxLen = 6
)

var (
	// ValidUsername 账号所能包含的字符串
	ValidUsername = regexp.MustCompile("^[a-zA-Z0-9-_]+$")

	// ValidPassword 账号所能包含的字符串
	ValidPassword = regexp.MustCompile("^[!-z]+$")
)

// Register 注册
func (userServiceImpl UserServiceImpl) Register(body request.UserRegisterBody) (response.UserRegister, error) {
	// 先检查是否已有这个账号:
	user, err := userServiceImpl.UserDao.GetUserByUsername(body.Username)
	if err == nil {
		// 账号重复
		return response.UserRegister{
			Response: response.FailWithMsg("账号重复"),
		}, err
	} else {
		// 不是因为找不到这个账号返回的异常
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.UserRegister{
				Response: response.Fail(),
			}, err
		} else {
			// 允许继续注册
		}
	}
	if !CheckUsername(body.Username) {
		return response.UserRegister{
			Response: response.FailWithMsg("账号长度过短或过长或包含了特殊字符"),
		}, err
	}
	if !CheckPassword(body.Password) {
		return response.UserRegister{
			Response: response.FailWithMsg("密码长度过短或过长或包含了特殊字符"),
		}, err
	}
	user = module.User{
		Username: body.Username,
		Password: body.Password,
	}
	mysqlDB := db.GetMysqlConnection()
	err = mysqlDB.Create(&user).Error
	if err != nil {
		return response.UserRegister{
			Response: response.FailWithMsg("注册失败, 请重试"),
		}, err
	}
	return response.UserRegister{
		Response: response.Success(),
		UserId:   0,
		Token:    "",
	}, err
}

// Login 登录
func (userServiceImpl UserServiceImpl) Login(body request.UserLoginBody) (response.UserLogin, error) {
	return response.UserLogin{}, nil
}

// Info 获取用户信息
func (userServiceImpl UserServiceImpl) Info(body request.UserInfoBody) (response.UserInfo, error) {
	return response.UserInfo{}, nil
}

// CheckUsername 检查账号是否合理
func CheckUsername(username string) bool {
	return len(username) >= UsernameMinLen && len(username) <= UsernameMaxLen && ValidUsername.Match([]byte(username))
}

// CheckPassword 检查密码是否合理
func CheckPassword(password string) bool {
	return len(password) >= PasswordMinLen && len(password) <= PasswordMaxLen && ValidPassword.Match([]byte(password))
}
