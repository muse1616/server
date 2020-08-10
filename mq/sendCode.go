package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"server/model"
	"server/utils"
)

// 单例连接
var conn, _ = GetRabbitMQConnection()

/**
代码发送到消息队列
*/
func SendCode(codeModel *model.CodeModel) (err error) {
	// 获取通道
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	//申明发送队列
	q, err := ch.QueueDeclare("code_submit", true, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare a queue")
	// 结构体转为字节数组
	body, err := json.Marshal(codeModel)
	utils.FailOnError(err, "Failed to marshal a struct")
	//	消息内容
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	utils.FailOnError(err, "Failed to publish a msg")
	return
}
