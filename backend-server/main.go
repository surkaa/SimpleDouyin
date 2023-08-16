package main

import (
	"SimpleDouyin/config"
	"SimpleDouyin/router"
	"SimpleDouyin/util"
	"github.com/gin-gonic/gin"
	"log"
)

const configFile = "../backend-server-configuration.json"

func main() {
	startServer()
}

// startServer 启动
func startServer() {
	g := gin.Default()
	initServer(g)
	if err := g.Run(); err != nil {
		log.Fatalf("启动失败")
	}
}

// initServer 初始化
func initServer(r *gin.Engine) {
	// 配置初始化
	config.LoadConfig(configFile)
	// 雪花id生成器初始化
	util.SF = util.NewSnowflake(config.Config.WorkerID)
	// 初始化日志打印器
	util.InitLogger("app.log")
	// 初始化路由
	router.RegisterRouter(r)
}
