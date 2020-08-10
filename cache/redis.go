package cache

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var Pool *redis.Pool

func InitRedis(server, pwd string) {
	if Pool == nil {
		Pool = NewPool(server, pwd)
	}
}

//redis 连接池
func NewPool(server, pwd string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if pwd != "" {
				//密码
				if _, err := c.Do("AUTH", pwd); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				log.Println("Redis连接成功")
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

/**
保存验证码至redis
*/
func SaveEmailVerification(email string, code string) (err error) {
	conn := Pool.Get()
	//ping
	err = Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return
	}
	//设置
	key := "email:verification:" + email
	// email:verification:1234@XX.com  1234
	_, err = conn.Do("Set", key, code)
	if err != nil {
		return
	}
	//3分钟过期
	_, err = conn.Do("expire", key, 180)
	if err != nil {
		return
	}
	return
}

/**
检查验证码
*/
func ConfirmEmailVerificationFromRedis(email string, verification string) (result bool, err error) {
	conn := Pool.Get()
	//ping
	err = Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return
	}
	// 此时因为内部为空 需要排除一个非空错误 此处不使用先判断key是否存在的情况
	correctVerification, err := redis.String(conn.Do("Get", "email:verification:"+email))
	if err == redis.ErrNil {
		result = false
		return result, nil
	}
	if correctVerification == verification {
		result = true
		// 删除redis中的key
		_, _ = conn.Do("DEL", "email:verification:"+email)
		return
	} else {
		result = false
		return
	}
}
