package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/cache"
	"server/utils"
)

//权限认证中间件
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//取cookie
		sessionId, err1 := c.Cookie(utils.Md5secret("session_id"))
		emaiMD5, err2 := c.Cookie(utils.Md5secret("email"))
		//根据session
		emailMD5Get, err3 := cache.GetSession(sessionId)
		if err1 != nil || err2 != nil || err3 != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"data":       "登录状态失效,请重新登录",
				"error_code": 5,
			})
			c.Abort()
			return
		}
		if emailMD5Get == emaiMD5 {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"data":       "登录状态失效,请重新登录",
				"error_code": 5,
			})
			c.Abort()
			return
		}

	}
}
