package impl

import (
	"SimpleDouyin/db"
	"SimpleDouyin/module"
	"github.com/sirupsen/logrus"
)

type UserDaoImpl struct {
	Logger *logrus.Logger
}

// 要使用impl时可以通过这样调用
// var userDaoImpl dao.UserDao = impl.UserDaoImpl{}

// InsertUser 向数据库中插入用户
func (userDao UserDaoImpl) InsertUser(user module.User) error {
	userDao.Logger.Info("InsertUser: user=%#v", user)
	return db.GetMysqlConnection().Create(&user).Error
}

// DeleteUserById 通过id删除用户 (将is_delete改为Deleted)
func (userDao UserDaoImpl) DeleteUserById(id int64) error {
	userDao.Logger.Info("DeleteUserById: id=%#v", id)
	return db.GetMysqlConnection().Model(&module.User{}).Where("Where id = ?", id).Updates(module.User{
		Entity: module.Entity{
			ID:       id,
			IsDelete: module.Deleted,
		},
	}).Error
}

// UpdateUserById 通过ID更新用户信息
func (userDao UserDaoImpl) UpdateUserById(user module.User) error {
	userDao.Logger.Info("UpdateUserById: user=%#v", user)
	return db.GetMysqlConnection().Model(&module.User{}).Where("where id = ?", user.ID).Updates(user).Error
}

// GetUserById 通过ID获取用户信息
func (userDao UserDaoImpl) GetUserById(id int64) (module.User, error) {
	userDao.Logger.Info("GetUserById: id=%#v", id)
	var user module.User
	result := db.GetMysqlConnection().First(&user, id)
	return user, result.Error
}

// GetUserByUsername 通过账号获取用户信息
func (userDao UserDaoImpl) GetUserByUsername(username string) (module.User, error) {
	userDao.Logger.Info("GetUserByUsername: username=%#v", username)
	var user module.User
	result := db.GetMysqlConnection().Where("username = ? AND is_delete = ?", username, module.BeforeDel).First(&user)
	return user, result.Error
}
