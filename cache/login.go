package cache

import (
	"github.com/garyburd/redigo/redis"
	"time"
)


/**
保存登录验证码至redis
*/
func SaveEmailVerificationForLogin(email string, code string) (err error) {
	conn := Pool.Get()
	//ping
	err = Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return
	}
	//设置
	key := "login:verification:" + email
	// email:verification:1234@XX.com  1234
	_, err = conn.Do("Set", key, code)
	if err != nil {
		return
	}
	//5分钟过期
	_, err = conn.Do("expire", key, 300)
	if err != nil {
		return
	}
	return
}

/**
检查登录验证码
 */
/**
检查验证码
*/
func ConfirmLoginEmailVerification(email string, verification string) (result bool, err error) {
	conn := Pool.Get()
	//ping
	err = Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return
	}
	// 此时因为内部为空 需要排除一个非空错误 此处不使用先判断key是否存在的情况
	correctVerification, err := redis.String(conn.Do("Get", "login:verification:"+email))
	if err == redis.ErrNil {
		result = false
		return result, nil
	}
	if correctVerification == verification {
		result = true
		// 删除redis中的key
		_, _ = conn.Do("DEL", "login:verification:"+email)
		return
	} else {
		result = false
		return
	}
}
