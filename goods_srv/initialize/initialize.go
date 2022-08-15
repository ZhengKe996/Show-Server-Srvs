package initialize

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"server_srvs/goods_srv/global"
)

// InitConfig 从配置文件中读取配置
func InitConfig() {
	configFileName := "goods_srv/config.yaml"
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.NaCosConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息: %v", &global.NaCosConfig)

	// 从 NaCos中读取配置信息
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: global.NaCosConfig.Host,
			Port:   global.NaCosConfig.Port,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         global.NaCosConfig.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NaCosConfig.DataId,
		Group:  global.NaCosConfig.Group})

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatalf("读取nacos配置失败： %s", err.Error())
	}
}
