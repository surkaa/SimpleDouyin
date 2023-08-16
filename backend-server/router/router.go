package router

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由 TODO 分组份文件注册
func RegisterRouter(r *gin.Engine) {
	if r == nil {
		return
	}
	registerAuthRouter(r.Group("/douyin/user/"))
}
