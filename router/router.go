package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/controller"
)

func SetupRouter() *gin.Engine {

	//默认路由
	r := gin.Default()

	//v1路由组
	v1Group := r.Group("v1")
	{
		//	用户提交代码
		v1Group.POST("/oj/code/submit", controller.CodeHandler)
	}

	//开启路由
	if err := r.Run("0.0.0.0:8888"); err != nil {
		log.Println(err)
		panic(err)
	}

	return r
}
