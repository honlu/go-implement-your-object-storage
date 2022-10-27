package heartbeat

import (
	"lib/rabbitmq"
	"os"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER")) // 创建一个rabbitmq.RabbitMQ结构体
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS")) // 向apiServers exchange交换机发送消息，把本服务节点的监听地址发送除去
		time.Sleep(5 * time.Second)                          // 每5s发送一次
	}
}
