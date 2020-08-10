package main

import (
	"log"
	"server/cache"
	"server/dao"
	"server/router"
	"server/utils"
)

func main() {
	// 配置文件
	var c utils.Config
	// 使用Yaml格式配置文件
	c = utils.YamlConfig{}
	m, err := c.LoadYamlConfig("conf")
	if err != nil {
		log.Fatalf("配置文件错误:%v", err)
		return
	}
	// redis连接池初始化
	cache.InitRedis(m["redis"]["address"].(string), m["redis"]["password"].(string))

	// mysql初始化
	err = dao.MysqlInit(m["mysql"]["username"].(string), m["mysql"]["password"].(string), m["mysql"]["address"].(string), m["mysql"]["dbName"].(string))
	if err != nil {
		log.Fatalf("MYSQL连接错误:%v", err)
		return
	}


	// 开启路由
	router.SetupRouter()

}
