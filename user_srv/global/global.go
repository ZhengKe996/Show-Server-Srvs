package global

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"server_srvs/user_srv/model"
	"strings"
	"time"
)

var (
	DB *gorm.DB
)

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = DB.AutoMigrate(&model.User{})
}

var options = &password.Options{16, 100, 32, sha512.New}

// Encryption 加密密码
func Encryption(code string) string {
	salt, encodedPwd := password.Encode(code, options)
	newpassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	return newpassword
}

// CheckPassWord 校验密码
func CheckPassWord(pwd, encryptedPwd string) bool {
	passwordInfo := strings.Split(encryptedPwd, "$")
	check := password.Verify(pwd, passwordInfo[2], passwordInfo[3], options)
	return check
}
