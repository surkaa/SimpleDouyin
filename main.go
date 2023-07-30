package main

import (
	"DouyinBackend/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.LoadConfig("./configuration.yaml")
	g := gin.Default()
	if err := g.Run(); err != nil {
		log.Fatalf("启动失败")
	}
}
