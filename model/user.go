package model

import "github.com/jinzhu/gorm"

// 用户注册模型 账号、邮箱、密码
type RegisterUserModel struct {
	Email string `json:"email"`
}

// 验证码验证
type VerificationConfirm struct {
	Email        string `json:"email"`
	Verification string `json:"Verification"`
	Password     string `json:"password"`
}

// 发送邮箱 以请求验证码
type EmailForVerification struct {
	Email string `json:"email"`
}

// 邮箱验证码登录
type LoginWithEmailVerification struct {
	Email        string `json:"email"`
	Verification string `json:"Verification"`
}

// 邮箱密码登录
type LoginWithPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthTest struct {
	Email string `json:"email"`
}

// 数据库User表
type User struct {
	gorm.Model
	// 邮箱 唯一性 非空
	Email    string
	Password string
	Nickname string
	RandomId int `gorm:"column:random_id"`
}
