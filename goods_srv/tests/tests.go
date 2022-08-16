package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"server_srvs/goods_srv/proto"
)

var Client proto.GoodsClient
var conn *grpc.ClientConn
var err error

func Init() {
	conn, err = grpc.Dial("127.0.0.1:50000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	Client = proto.NewGoodsClient(conn)
}

func TestGetBrandsList() {
	list, err := Client.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(list.Total)
	for _, brands := range list.Data {
		fmt.Println(brands.Name, brands.Logo)
	}
}

func TestCreateBrands() {
	for i := 0; i < 10; i++ {
		req, err := Client.CreateBrand(context.Background(), &proto.BrandRequest{
			Name: fmt.Sprintf("Timu:%d", i),
			Logo: "https://front-zk.oss-cn-hangzhou.aliyuncs.com/user.jpg",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(req.Id)
	}
}

func TestUpdateBrands() {
	_, err := Client.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   2,
		Name: "Baidu",
		Logo: "https://front-zk.oss-cn-hangzhou.aliyuncs.com/logo.png",
	})
	if err != nil {
		panic(err)
	}
}

func TestDeleteBrands() {
	_, err := Client.DeleteBrand(context.Background(), &proto.BrandRequest{Id: 2})
	if err != nil {
		panic(err)
	}
}
func TestGetBanner() {
	list, err := Client.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(list.Total)
	for _, value := range list.Data {
		fmt.Println(value)
	}
}
func main() {
	Init()
	//TestCreateBrands()
	//TestGetBrandsList()
	//TestUpdateBrands()
	//TestDeleteBrands()
	//TestGetBanner()
	defer conn.Close()
}
