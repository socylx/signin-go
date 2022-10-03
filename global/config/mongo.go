package config

type MongoConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Auth     string
	Database string
}

var Mongo = &MongoConfig{}
