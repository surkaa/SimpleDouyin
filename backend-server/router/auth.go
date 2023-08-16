package router

import (
	"SimpleDouyin/controller"
	"github.com/gin-gonic/gin"
)

// registerAuthRouter [登陆, 注册, 用户信息]路由
func registerAuthRouter(d *gin.RouterGroup) {
	if d == nil {
		return
	}
	d.POST("register/", controller.Register)
	d.POST("login/", controller.Login)
	d.GET("", controller.Info)
}
