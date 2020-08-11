package dao

import (
	"server/model"
	"server/utils"
)

func UserLoginWithPassword(email string, password string) bool {
	var user model.User
	isCorrect := DB.Table("user").Debug().Where("`email` = ? and `password` = ?", email, utils.Md5secret(utils.Md5secret(password))).First(&user).RecordNotFound()
	return !isCorrect
}
