package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	g := gin.Default()
	if err := g.Run(); err != nil {
		log.Fatalf("启动失败")
	}
}
