package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null comment '用户手机号'"`
	Password string     `gorm:"type:varchar(100);not null  comment '用户密码(加密)'"`
	NikeName string     `gorm:"type:varchar(20)  comment '用户名 默认为null'"`
	Birthday *time.Time `gorm:"type:datetime  comment '用户生日'"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female 表示女,male 表示男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1 表示普通用户 0 表示管理员'"`
}
