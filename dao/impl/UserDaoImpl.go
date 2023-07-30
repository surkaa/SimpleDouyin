package impl

import (
	"DouyinBackend/module"
)

type UserDaoImpl struct {
}

// 要使用impl时可以通过这样调用
// var userDaoImpl dao.UserDao = UserDaoImpl{}

// InsertUser 向数据库中插入用户
func (userDao UserDaoImpl) InsertUser(user module.User) error {
	return nil
}

// DeleteUserById 通过id删除用户 (将is_delete改为1)
func (userDao UserDaoImpl) DeleteUserById(id int64) error {
	return nil
}

// UpdateUserById 通过ID更新用户信息
func (userDao UserDaoImpl) UpdateUserById(id int64) error {
	return nil
}

// GetUserById 通过ID获取用户信息
func (userDao UserDaoImpl) GetUserById(id int64) (module.User, error) {
	return module.User{}, nil
}
