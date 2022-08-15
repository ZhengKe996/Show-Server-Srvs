package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"server_srvs/user_srv/global"
	"server_srvs/user_srv/handler"
	"server_srvs/user_srv/initialize"
	"server_srvs/user_srv/proto"
	"server_srvs/user_srv/utils"
	"syscall"
)

func main() {
	// 初始化
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()

	// 命令解析
	IP := flag.String("ip", "0.0.0.0", "IP地址")
	Port := flag.Int("port", 0, "IP地址")

	flag.Parse()
	if port, err := utils.GetFreePort(); *Port == 0 && err == nil {
		*Port = port
	}
	zap.S().Infof("IP:%s, Port:%d", *IP, *Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成对应的检查对象
	check := new(api.AgentServiceCheck)
	check.GRPC = fmt.Sprintf("192.168.0.162:%d", *Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "10s"

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	registration.ID = serviceID
	registration.Port = *Port
	registration.Tags = []string{"user", "srv"}
	registration.Address = "192.168.0.162"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	go func() {
		err = server.Serve(listen)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	// 优雅退出：接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serviceID); err != nil {
		zap.S().Info("【UserSrv服务】注销失败")
	}
	zap.S().Info("【UserSrv服务】注销成功")
}
