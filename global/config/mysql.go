package config

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

var Mysql = &MysqlConfig{}
