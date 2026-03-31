package rabbitmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yuanji6666/gopherAI/config"
)

// conn 全局连接实例，应用程序和RabbitMQ之间的TCP连接
// 一个连接可以承载多个通道
var conn *amqp.Connection

// InitConn 初始化连接
func InitConn(){
	mqURL := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		config.GetConfig().RabbitmqUsername,
		config.GetConfig().RabbitmqPassword,
		config.GetConfig().RabbitmqHost,
		config.GetConfig().RabbitmqPort,
		config.GetConfig().RabbitmqVhost,
	)

	log.Println("mqURL :", mqURL)

	var err error
	conn, err = amqp.Dial(mqURL)

	if err != nil {
		log.Fatalf("RabbitMQ connection failed: %v", err) // 输出错误并退出程序
	}
}

// RabbitMQ 封装连接，通道，交换机，路由键
// 连接：应用程序和RabbitMQ之间的TCP连接，一个连接可以承载多个通道
// 通道：在连接上的虚拟通信层，允许多个通道复用一个TCP连接，提高效率和并发性。
// 交换机：接受消息和路由键，根据绑定规则将消息路由到响应队列
// 路由键：标签字符串，交换机使用路由键匹配队列绑定，决定消息去向。在direct交换机中，路由键必须精确匹配绑定键。
type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	Exchange string
	Key      string
}

// NewRabbitMQ 构造exchange，key字段
func NewRabbitMQ(exchange, key string) *RabbitMQ {
	return &RabbitMQ{Exchange: exchange, Key: key}
}

// NewWorkRabbitMQ 以direct路由模式，queue队列名初始化（队列还未初始化）
func NewWorkRabbitMQ(queue string) *RabbitMQ{
	if conn == nil {
		InitConn()
	}

	rabbitMQ := NewRabbitMQ("", queue)
	rabbitMQ.conn = conn 

	var err error
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()

	if err != nil{
		panic(err.Error())
	}

	return rabbitMQ
}

// Destory 关闭连接和通道
func (r *RabbitMQ)Destory(){
	r.channel.Close()
	r.conn.Close()
}

// Publish 发送msg
func (r *RabbitMQ) Publish(msg []byte) error {
	_, err := r.channel.QueueDeclare(
		r.Key, 
		false,
		false,
		false,
		false,
		nil,	
	)

	if err != nil {
		return err
	}

	return r.channel.Publish(r.Exchange,r.Key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body: msg,
	})

}

// Consume 接收消息，调用handle方法处理
func (r *RabbitMQ) Consume(handle func(msg *amqp.Delivery) error) {
	q, err := r.channel.QueueDeclare(
		r.Key, 
		false,
		false,
		false,
		false,
		nil,	
	)

	if err != nil {
		panic(err)
	}

	msgs, err := r.channel.Consume(q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	for msg := range msgs{
		err := handle(&msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}