package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/controller"
	"server/middleware"
)

func SetupRouter() *gin.Engine {
	//默认路由
	r := gin.Default()
	//v1路由组
	v1Group := r.Group("v1")
	{

		//  测试路由
		v1Group.POST("/test", controller.TestHandler)
		// 	测试登录状态授权
		v1Group.POST("/auth/login", middleware.LoginAuth(), controller.AuthHandler)
		//	用户提交代码
		v1Group.POST("/code/submit", controller.CodeHandler)
		//  用户注册
		v1Group.POST("/user/register", controller.UserRegisterController)
		//	注册验证码确认
		v1Group.POST("/user/verification_confirm", controller.VerificationConfirm)
		//	请求邮箱验证码登录
		v1Group.POST("/user/login/request/verification", controller.SendLoginEmailVerification)
		//  验证登录验证码
		v1Group.POST("/user/login/confirm/verification", controller.ConfirmLoginEmailVerification)
		//  邮箱密码登录
		v1Group.POST("/user/login/password",controller.LoginWithPassword)

	}

	//开启路由
	if err := r.Run("0.0.0.0:8888"); err != nil {
		log.Println(err)
		panic(err)
	}

	return r
}
