package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"server_srvs//global"
	"server_srvs//model"
	"server_srvs//proto"
	"time"
)

type UserServer struct {
	proto.UnimplementedUserServer
}

// Model2Response 模型转UserInfoResponse
func Model2Response(user model.User) proto.UserInfoResponse {
	// 在grpc的message中字段有默认值，不能随便赋值nil进去，容易出错
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.Password,
		Mobile:   user.Mobile,
		NickName: user.NikeName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

// Paginate 分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// GetUserList 获取用户列表
func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	rsp := &proto.UserListResponse{}
	// 查询总数
	rsp.Total = int32(result.RowsAffected)
	// 分页数据
	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	// 处理数据
	for _, user := range users {
		userInfoResp := Model2Response(user)
		rsp.Data = append(rsp.Data, &userInfoResp)
	}

	return rsp, nil
}

// GetUserByMobile 通过手机号查询用户
func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := Model2Response(user)
	return &userInfoRsp, nil
}

// GetUserById 通过ID查询用户
func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := Model2Response(user)
	return &userInfoRsp, nil
}

// CreateUser 新建用户
func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Error(codes.AlreadyExists, "用户已存在")
	}
	user.Mobile = req.Mobile
	user.NikeName = req.NickName

	// 密码加密
	user.Password = global.Encryption(req.PassWord)

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	userInfoRsp := Model2Response(user)
	return &userInfoRsp, nil
}

// UpdateUser 个人中心更新用户
func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	birthday := time.Unix(int64(req.BirthDay), 0)
	user.NikeName = req.NickName
	user.Birthday = &birthday
	user.Gender = req.Gender

	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

// CheckPassWord 校验密码
func (s *UserServer) CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	check := global.CheckPassWord(req.Password, req.EncryptedPassword)
	return &proto.CheckResponse{Success: check}, nil
}
