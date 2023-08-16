package dao

import (
	"SimpleDouyin/dao/impl"
	"SimpleDouyin/module"
	"github.com/sirupsen/logrus"
)

// NewUserDao UserDao 实例化工厂函数
func NewUserDao() UserDao {
	return &impl.UserDaoImpl{
		Logger: logrus.New(),
	}
}

type UserDao interface {
	// InsertUser 向数据库中插入用户
	InsertUser(user module.User) error

	// DeleteUserById 通过id删除用户 (将is_delete改为1)
	DeleteUserById(id int64) error

	// UpdateUserById 通过ID更新用户信息
	UpdateUserById(user module.User) error

	// GetUserById 通过ID获取用户信息
	GetUserById(id int64) (module.User, error)

	// GetUserByUsername 通过账号获取用户信息
	GetUserByUsername(username string) (module.User, error)
}
