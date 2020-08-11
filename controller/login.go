package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"server/cache"
	"server/dao"
	"server/model"
	"server/utils"
	"strconv"
	"time"
)

/**
发送邮箱验证码登录
*/
func SendLoginEmailVerification(ctx *gin.Context) {
	var efv model.EmailForVerification
	// 参数错误
	if err := ctx.Bind(&efv); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "参数错误",
		})
		return
	}
	// 验证邮箱是否已注册
	isRegistered := dao.IsEmailRegistered(efv.Email)
	if isRegistered == false {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "该邮箱未注册",
		})
		return
	}

	// 发送邮箱验证码
	// 生成验证码 随机四位数
	rand.Seed(time.Now().Unix())
	verification := strconv.Itoa(rand.Intn(9000) + 1000)
	if err := utils.SendEmailVerificationForLogin(efv.Email, verification); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 2,
			"data":       "验证码发送失败",
		})
		log.Println(err)
		return
	}
	// 2. 将验证码存入redis 有效时间5分钟
	if err := cache.SaveEmailVerificationForLogin(efv.Email, verification); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 3,
			"data":       "服务器错误",
		})
		log.Println(err)
		return
	}
	// 3. 显示成功发送
	ctx.JSON(http.StatusOK, gin.H{
		"error_code": 0,
		"data":       "验证码已发送",
	})

}

/**
验证登录邮箱验证码
email verification
*/
func ConfirmLoginEmailVerification(ctx *gin.Context) {
	var loginWithEmailVerification model.LoginWithEmailVerification
	// 参数错误
	if err := ctx.Bind(&loginWithEmailVerification); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "参数错误",
		})
		return
	}
	//	检查登录验证码是否正确
	result, err := cache.ConfirmLoginEmailVerification(loginWithEmailVerification.Email, loginWithEmailVerification.Verification)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 3,
			"data":       "服务器错误",
		})
		log.Println(err)
		return
	}
	// 验证码存在且正确 返回true
	if result == true {
		sessionId := utils.GenerateUUid()
		// 处理session_id
		err := cache.SetSession(sessionId, loginWithEmailVerification.Email)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error_code": 3,
				"data":       "服务器错误",
			})
			log.Println(err)
			return
		}
		// 将session_id保存在本地的cookie
		ctx.SetCookie(utils.Md5secret("session_id"), sessionId, 3600*12, "/", "127.0.0.1", false, true)
		// 将email加密放在本地cookie md5
		ctx.SetCookie(utils.Md5secret("email"), utils.Md5secret(loginWithEmailVerification.Email), 3600*24, "/", "127.0.0.1", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": 0,
			"data":       "登录成功",
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "验证码错误或已过期",
		})
	}
}

/**
使用密码登录
*/
func LoginWithPassword(ctx *gin.Context) {
	var lwp model.LoginWithPassword
	if err := ctx.Bind(&lwp); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "参数错误",
		})
	}
	result := dao.UserLoginWithPassword(lwp.Email, lwp.Password)
	if result == true {
		sessionId := utils.GenerateUUid()
		// 处理session_id
		err := cache.SetSession(sessionId, lwp.Email)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error_code": 3,
				"data":       "服务器错误",
			})
			log.Println(err)
			return
		}
		// 将session_id保存在本地的cookie
		ctx.SetCookie(utils.Md5secret("session_id"), sessionId, 3600*12, "/", "127.0.0.1", false, true)
		// 将email加密放在本地cookie md5
		ctx.SetCookie(utils.Md5secret("email"), utils.Md5secret(lwp.Email), 3600*24, "/", "127.0.0.1", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": 0,
			"data":       "登录成功",
		})
	} else if result == false {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "账号或密码错误",
		})
	}
}
