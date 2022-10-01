package config

import "fmt"

type ServerConfig struct {
	Mode          string
	Port          string
	Token         string
	ServerLogFile string
}

var Server *ServerConfig = &ServerConfig{}

func (server *ServerConfig) String() string {
	return fmt.Sprintf("Server: %s, Port: %v", server.Mode, server.Port)
}
