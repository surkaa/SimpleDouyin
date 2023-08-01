package router

import (
	"DouyinBackend/controller"
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由 TODO 分组份文件注册
func RegisterRouter(r *gin.Engine) {
	d := r.Group("/douyin")
	d.POST("/user/register/", controller.Register)
	d.POST("/user/login", controller.Login)
	d.GET("/user/", controller.Info)
}
