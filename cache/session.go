package cache

import (
	"github.com/garyburd/redigo/redis"
	"server/utils"
	"time"
)

//设置redis中session Key--uuid Value--username
func SetSession(sessionId, email string) (err error) {
	conn := Pool.Get()
	//ping
	err = Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return err
	}
	//设置redis session
	_, err = conn.Do("Set", "session:"+sessionId, utils.Md5secret(email))
	if err != nil {
		return err
	}
	// 12小时过期
	_, err = conn.Do("expire", "session:"+sessionId, 12*3600)
	if err != nil {
		return err
	}
	return
}

//根据sessionId在redis中取值
func GetSession(sessionId string) (emailMD5 string, err error) {
	conn := Pool.Get()
	//ping
	err = Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return
	}
	emailMD5, err = redis.String(conn.Do("Get", "session:"+sessionId))
	if err != nil {
		return
	}
	return
}
