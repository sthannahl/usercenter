// +build ignore
package crypto

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncryptByBcrypt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(hash)

	encodePW := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println(encodePW)
	return encodePW
}

func CompareHashAndPassword(pwd, encodePW string) bool {
	// 正确密码验证
	err := bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(pwd))
	return err == nil
}

func main() {
	CompareHashAndPassword("123456", EncryptByBcrypt("123456"))
}
