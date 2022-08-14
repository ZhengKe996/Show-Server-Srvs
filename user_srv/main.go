package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"server_srvs/user_srv/handler"
	"server_srvs/user_srv/initialize"
	"server_srvs/user_srv/proto"
)

func main() {
	// 初始化
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()

	// 命令解析
	IP := flag.String("ip", "0.0.0.0", "IP地址")
	Port := flag.Int("port", 50051, "IP地址")

	flag.Parse()
	zap.S().Infof("IP:%s, Port:%d", *IP, *Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(listen)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
