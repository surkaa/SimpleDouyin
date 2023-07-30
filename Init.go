package main

import (
	"DouyinBackend/config"
	"DouyinBackend/util"
)

func initApp() {
	config.LoadConfig("./configuration.json")
	util.SF = util.NewSnowflake(config.Config.WorkerID)
}
