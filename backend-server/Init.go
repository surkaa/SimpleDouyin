package main

import (
	"DouyinBackend/config"
	"DouyinBackend/router"
	"DouyinBackend/util"
	"github.com/gin-gonic/gin"
)

func initApp(r *gin.Engine) {
	config.LoadConfig("../backend-server-configuration.json")
	util.SF = util.NewSnowflake(config.Config.WorkerID)
	util.InitLogger("app.log")
	router.RegisterRouter(r)
}
