package db

import (
	"DouyinBackend/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetMysqlConnection 获取Mysql连接
func GetMysqlConnection() (*gorm.DB, error) {
	// 读取到之后转成string
	db, err := gorm.Open(mysql.Open(config.Config.MysqlDNS), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
