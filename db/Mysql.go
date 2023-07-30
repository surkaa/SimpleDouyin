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
	dns, ok := config.Config["mysqlDNS"].(string)
	if !ok {
		return nil, fmt.Errorf("无法获取mysqlDNS(string)配置")
	}
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
