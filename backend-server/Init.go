package main

import (
	"DouyinBackend/config"
	"DouyinBackend/router"
	"DouyinBackend/util"
	"github.com/gin-gonic/gin"
	"log"
)

func startServer() {
	g := gin.Default()
	initApp(g)
	if err := g.Run(); err != nil {
		log.Fatalf("启动失败")
	}
}

func initApp(r *gin.Engine) {
	config.LoadConfig("../backend-server-configuration.json")
	util.SF = util.NewSnowflake(config.Config.WorkerID)
	util.InitLogger("app.log")
	router.RegisterRouter(r)
}
