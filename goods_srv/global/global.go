package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"server_srvs/goods_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NaCosConfig  config.NaCosConfig
	EsClient     *elastic.Client
)
