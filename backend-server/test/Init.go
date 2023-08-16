package test

import (
	"SimpleDouyin/config"
	"SimpleDouyin/util"
)

func initTestApp() {
	config.LoadConfig("./configuration.json")
	util.SF = util.NewSnowflake(config.Config.WorkerID)
	util.InitLogger("app.log")
}
