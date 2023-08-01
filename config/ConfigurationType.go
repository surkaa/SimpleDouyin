package config

type Configuration struct {
	MysqlDNS string `json:"mysqlDNS"`
	WorkerID int64  `json:"workerID"`
	Redis    Redis  `json:"redis"`
}

// Redis Redis配置
type Redis struct {
	Addr     string `json:"Addr"`     // 链接地址及其端口
	Password string `json:"Password"` // no password set
	DB       int    `json:"DB"`       // use default DB
	PoolSize int    `json:"PoolSize"` // 连接池大小
}
