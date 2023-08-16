package test

import (
	"DouyinBackend/config"
	"DouyinBackend/util"
)

func initTestApp() {
	config.LoadConfig("./configuration.json")
	util.SF = util.NewSnowflake(config.Config.WorkerID)
	util.InitLogger("app.log")
}
