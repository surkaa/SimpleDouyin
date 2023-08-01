package controller

import (
	"DouyinBackend/request"
	"DouyinBackend/response"
	"DouyinBackend/service"
	"DouyinBackend/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userService service.UserService = impl.UserServiceImpl{}

// Register 注册
func Register(c *gin.Context) {
	registerBody, err := request.ParseUserRegisterBody(c)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	registerResponse, err := userService.Register(registerBody)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, registerResponse)
}

// Login 登录
func Login(c *gin.Context) {
	loginBody, err := request.ParseUserLoginBody(c)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	loginResponse, err := userService.Login(loginBody)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, loginResponse)
}

// Info 获取用户信息
func Info(c *gin.Context) {
	infoBody, err := request.ParseUserInfoBody(c)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	infoResponse, err := userService.Info(infoBody)
	if err != nil {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, infoResponse)
}
