package utils

import uuid "github.com/satori/go.uuid"

//生成uuid
func GenerateUUid() string {
	// 时间戳 加 MAC地址 待修改
	return uuid.NewV1().String()
}