package dao

import (
	"server/model"
	"server/utils"
)

/**
用户邮箱是否已注册
*/
func IsEmailRegistered(email string) (isRegistered bool) {
	var user model.User
	//DB.Table("user").Debug().Where("`email` = ?", email).First(&user)
	//fmt.Printf("%+v\n\n", user)
	// 记录未找到 返回false 即为注册 可注册
	isRegistered = DB.Table("user").Debug().Where("`email` = ?", email).First(&user).RecordNotFound()
	// 没找到isRegistered应该true
	return !isRegistered
}

/**
注册用户 密码使用md5*2
*/
func UserRegister(email string, password string) error {
	password = utils.Md5secret(utils.Md5secret(password))
	//	写入数据库
	user := &model.User{
		Email:    email,
		Password: password,
	}
	err := DB.Table("user").Create(&user).Error
	return err
}
