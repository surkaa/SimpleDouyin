package test

import (
	"DouyinBackend/config"
	"DouyinBackend/dao"
	"DouyinBackend/dao/impl"
	"fmt"
	"testing"
)

// TestGetUserByUsername 测试userDao.GetUserByUsername方法
func TestGetUserByUsername(t *testing.T) {
	config.LoadConfig("../configuration.json")
	var userDao dao.UserDao = impl.UserDaoImpl{}
	user, err := userDao.GetUserByUsername("test")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("User = %#v", user)
}
