package router

import (
	"SimpleDouyin/controller"
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由 TODO 分组份文件注册
func RegisterRouter(r *gin.Engine) {
	if r == nil {
		return
	}
	d := r.Group("/douyin")
	d.POST("/user/register/", controller.Register)
	d.POST("/user/login/", controller.Login)
	d.GET("/user/", controller.Info)
}
