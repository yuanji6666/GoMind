// Package rabbitmq 基于RabbitMQ的消息队列系统
package rabbitmq

var (
	RMQMessage *RabbitMQ

)

func InitRabbitMQ(){
	RMQMessage = NewWorkRabbitMQ("message")
	go RMQMessage.Consume(HandleMQMessage) //异步运行，不阻塞InitRabbitMQ
}

func DestoryRabbitMQ(){
	RMQMessage.Destory()
}