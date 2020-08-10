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
		//  测试路由
		v1Group.POST("/test", controller.TestHandler)
		//	用户提交代码
		v1Group.POST("/code/submit", controller.CodeHandler)
		//  用户注册
		v1Group.POST("/user/register", controller.UserRegisterController)
		//	注册验证码确认
		v1Group.POST("/user/verification_confirm",controller.VerificationConfirm)
	}

	//开启路由
	if err := r.Run("0.0.0.0:8888"); err != nil {
		log.Println(err)
		panic(err)
	}

	return r
}
