package global

import (
	"gorm.io/gorm"
	"server_srvs/goods_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NaCosConfig  config.NaCosConfig
)
