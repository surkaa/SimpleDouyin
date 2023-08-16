package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// RequestInfoMiddleware 打印每次请求的请求信息
func RequestInfoMiddleware(c *gin.Context) {
	// 获取请求方法、请求路径和请求 IP
	method := c.Request.Method
	path := c.Request.URL.Path
	ip := c.ClientIP()

	// 获取查询参数
	query := c.Request.URL.Query()

	// 将请求信息打印出来
	fmt.Printf("收到请求 - Method: %s, Path: %s, Query: %v, IP: %s\n", method, path, query, ip)

	// 继续处理请求
	c.Next()
}
