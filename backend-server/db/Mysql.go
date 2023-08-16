package db

import (
	"DouyinBackend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// GetMysqlConnection 获取Mysql连接
func GetMysqlConnection() *gorm.DB {
	// 读取到之后转成string
	db, err := gorm.Open(mysql.Open(config.Config.MysqlDNS), &gorm.Config{})
	if err != nil {
		log.Fatalf("获取数据库连接失败: %v", err)
	}
	return db
}
