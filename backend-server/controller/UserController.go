package controller

import (
	"SimpleDouyin/request"
	"SimpleDouyin/response"
	"SimpleDouyin/service"
	"SimpleDouyin/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var userService = service.NewUserService()
var logger = logrus.New()

// Register 注册
func Register(c *gin.Context) {
	util.Log().WithTime(time.Now().Local()).Info("收到注册请求\t")
	registerBody, err := request.ParseUserRegisterBody(c)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	logger.Info("解析注册体: \t%#v\n", registerBody)
	registerResponse, err := userService.Register(registerBody)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	logger.Info("注册结果: %v\n", registerResponse)
	c.JSON(http.StatusOK, registerResponse)
}

// Login 登录
func Login(c *gin.Context) {
	logger.Info("收到登录请求\t")
	loginBody, err := request.ParseUserLoginBody(c)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	logger.Info("解析登录体: %v\t", loginBody)
	loginResponse, err := userService.Login(loginBody)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	logger.Info("登录结果: %v\n", loginResponse)
	c.JSON(http.StatusOK, loginResponse)
}

// Info 获取用户信息
func Info(c *gin.Context) {
	logger.Info("收到获取用户信息请求\t")
	infoBody, err := request.ParseUserInfoBody(c)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	logger.Info("解析获取用户信息体: %v\t", infoBody)
	infoResponse, err := userService.Info(infoBody)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	logger.Info("获取用户信息结果: %v\n", infoResponse)
	c.JSON(http.StatusOK, infoResponse)
}
