package main

import (
	"DouyinBackend/config"
	"DouyinBackend/router"
	"DouyinBackend/util"
	"github.com/gin-gonic/gin"
)

func initApp(r *gin.Engine) {
	config.LoadConfig("./configuration.json")
	util.SF = util.NewSnowflake(config.Config.WorkerID)
	router.RegisterRouter(r)
}
