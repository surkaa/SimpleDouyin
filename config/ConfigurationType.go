package config

type Configuration struct {
	MysqlDNS string `json:"mysqlDNS"`
	WorkerID int64  `json:"workerID"`
}
