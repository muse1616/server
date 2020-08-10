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
注册路由处理句柄
//	1.邮箱
*/
func UserRegisterController(ctx *gin.Context) {
	// bind Model
	var registerUserModel model.RegisterUserModel
	// 参数错误
	if err := ctx.Bind(&registerUserModel); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "参数错误",
		})
		return
	}
	// ** 检查该邮箱是否注册 若已注册则告知已注册
	isRegister := dao.IsEmailRegistered(registerUserModel.Email)

	if isRegister == true {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "该邮箱已注册",
		})
		return
	}
	// 注册流程
	// 1. 邮箱验证码发送 此处注意原则 就算邮箱错误也无需告诉用户 直接存入redis即可 前端注意发送时间控制 1分钟内不要重复发送
	// 生成验证码 随机四位数
	rand.Seed(time.Now().Unix())
	verification := strconv.Itoa(rand.Intn(9000) + 1000)
	if err := utils.SendEmailVerification(registerUserModel.Email, verification); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 2,
			"data":       "验证码发送失败",
		})
		log.Println(err)
		return
	}
	// 2. 将验证码存入redis 有效时间30分钟
	if err := cache.SaveEmailVerification(registerUserModel.Email, verification); err != nil {
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
注册验证邮箱验证码
*/
func VerificationConfirm(ctx *gin.Context) {
	var vConfirm model.VerificationConfirm
	// 参数错误
	if err := ctx.Bind(&vConfirm); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "参数错误",
		})
		return
	}
	r, err := cache.ConfirmEmailVerificationFromRedis(vConfirm.Email, vConfirm.Verification)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 3,
			"data":       "服务器错误",
		})
		log.Println(err)
		return
	}
	// 验证码存在且正确 返回true
	if r == true {
		// 数据库注册用户
		err := dao.UserRegister(vConfirm.Email, vConfirm.Password)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error_code": 3,
				"data":       "服务器错误",
			})
			log.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": 0,
			"data":       "邮箱: " + vConfirm.Email + " 注册成功",
		})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error_code": 1,
			"data":       "验证码错误或已过期",
		})
	}
}
