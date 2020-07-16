package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	"server/mq"
)

func CodeHandler(ctx *gin.Context) {
	//	代码提交模型
	var codeModel model.CodeModel
	if err := ctx.Bind(&codeModel); err != nil {
		//	Bind JSON
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "fail",
		})
		return
	}

	//发送代码包到消息队列等待处理
	err := mq.SendCode(&codeModel)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": "fail",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}
