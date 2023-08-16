package test

import (
	"SimpleDouyin/dao"
	"fmt"
	"testing"
)

// TestGetUserByUsername 测试userDao.GetUserByUsername方法
func TestGetUserByUsername(t *testing.T) {
	initTestApp()
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByUsername("test")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("User = %#v", user)
}
