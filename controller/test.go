package controller

import (
	"github.com/gin-gonic/gin"
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
