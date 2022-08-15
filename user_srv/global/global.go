package global

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/gorm"
	"server_srvs/user_srv/config"
	"strings"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NaCosConfig  config.NaCosConfig
)

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
