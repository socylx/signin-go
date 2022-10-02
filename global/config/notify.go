package config

type NotifyConfig struct {
	Host string
	Port int
	User string
	Pass string
	To   string
}

var Notify = &NotifyConfig{}
