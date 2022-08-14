package initialize

import (
	"github.com/spf13/viper"
	"server_srvs/user_srv/global"
)

// InitConfig 从配置文件中读取配置
func InitConfig() {
	configFileName := "user_srv/config.yaml"
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
	//zap.S().Infof("配置信息: %v", global.ServerConfig)

}
