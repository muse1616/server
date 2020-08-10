package model

/**
用户提交代码结构体 以json格式提交
json:
{
	"id":"用户账号",
	"code":"代码段",
	"type":"语言类型",
	"pId":"题目编号",
}
*/
type CodeModel struct {
	Id   string `json:"id"`
	Code string `json:"code"`
	Type string `json:"type"`
	PId  string `json:"pId"`
}
