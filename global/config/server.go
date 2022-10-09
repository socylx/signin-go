package config

import "fmt"

type ServerConfig struct {
	Name                 string
	Mode                 string
	Port                 string
	MaxRequestsPerSecond int
	Token                string
	ServerLogDir         string
	ServerLogFile        string
	SQLLogFile           string

	IsRelease bool
}

var Server *ServerConfig = &ServerConfig{}

func (server *ServerConfig) String() string {
	return fmt.Sprintf("Server: %s, Port: %v", server.Mode, server.Port)
}
