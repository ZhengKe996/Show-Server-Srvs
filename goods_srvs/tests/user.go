package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"server_srvs//global"
	"server_srvs//proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn
var err error

func Init() {
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

// TestGetUserList 测试查询用户接口
func TestGetUserList() {
	list, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range list.Data {
		fmt.Println(user.Mobile, user.NickName, user.PassWord)

		check, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          "admin123",
			EncryptedPassword: user.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(check.Success)
	}
}

// TestCreateUser 测试创建用户接口
func TestCreateUser() {
	encryption := "admin123"
	for i := 0; i < 10; i++ {
		req, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("zhangsan%d", i),
			Mobile:   fmt.Sprintf("1981689629%d", i),
			PassWord: encryption,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(req.PassWord)
	}
}

// TestGetUserByID 测试ID查询
func TestGetUserByID() {
	for i := 21; i < 31; i++ {
		user, err := userClient.GetUserById(context.Background(), &proto.IdRequest{Id: int32(i)})
		if err != nil {
			panic(err)
		}
		fmt.Println(user.Mobile, user.NickName)

	}
}

// TestGetUserByMobile 测试手机号查询
func TestGetUserByMobile() {
	for i := 0; i < 10; i++ {
		user, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: fmt.Sprintf("1981689629%d", i)})
		if err != nil {
			panic(err)
		}
		fmt.Println(user.Id, user.NickName)
	}
}

func TestUpdateUser() {
	for i := 21; i < 31; i++ {
		_, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
			Id:       int32(i),
			NickName: fmt.Sprintf("lisi%d", i),
			BirthDay: 0,
		})

		if err != nil {
			panic(err)
		}
	}
}

func TestCheckPassword() {
	word := global.CheckPassWord("admin123", "$pbkdf2-sha512$TQLXmldWtgJTQsCZ$fba4ab8b2c3f190aa9af7c630622de9b0796491c748d2f8192134cf2c965557a")
	fmt.Println(word)
}
func main() {
	Init()
	TestGetUserList()
	//TestCreateUser()
	//TestGetUserByID()
	//TestGetUserByMobile()
	//TestUpdateUser()
	//TestCheckPassword()
	defer conn.Close()
}
