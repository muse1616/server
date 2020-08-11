package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/model"
)

/**
测试路由
*/
func TestHandler(ctx *gin.Context) {
	//	代码提交模型
	var testModel model.TestModel
	if err := ctx.ShouldBind(&testModel); err != nil {
		//	Bind JSON
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "fail",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"id":     testModel.TestID,
	})
}

// 测试登录授权
func AuthHandler(ctx *gin.Context) {
	//测试授权状态
	var authTest model.AuthTest
	if err := ctx.ShouldBind(&authTest); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "参数错误",
		})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":       "登录状态有效,账号:" + authTest.Email,
		"error_code": 0,
	})
}
