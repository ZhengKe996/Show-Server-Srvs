package global

import (
	"gorm.io/gorm"
	"server_srvs//config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NaCosConfig  config.NaCosConfig
)
