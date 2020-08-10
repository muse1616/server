package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"server/utils"
)

/**
	RabbitMQ 消息队列连接
	配置参数:
	id: RabbitMQ用户
	pwd:RabbitMQ密码
	URL:RabbitMQ URL
	port:端口号
	vhost:虚拟主机
 */
func GetRabbitMQConnection() (conn *amqp.Connection, err error) {
	// 配置文件
	var c utils.Config
	// 使用Yaml格式配置文件
	c = utils.YamlConfig{}
	content, err := c.LoadYamlConfig("conf")
	if err != nil {
		return
	}
	// url
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/%s",
		content["code"]["id"].(string),
		content["code"]["pwd"].(string),
		content["code"]["url"].(string),
		content["code"]["port"].(string),
		content["code"]["vhost"].(string),
	)
	// 获取连接
	conn, err = amqp.Dial(url)
	if err != nil {
		log.Println(err)
	}
	// 返回连接
	return
}
