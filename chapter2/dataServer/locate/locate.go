package locate

import (
	"lib/rabbitmq"
	"os"
	"strconv"
)

// 用于定位对象
func Locate(name string) bool {
	_, err := os.Stat(name) // os.Stat访问磁盘对应的文件名
	return !os.IsNotExist(err)
}

// 用于监听定位消息
func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers") // 绑定dataServers exchange.
	c := q.Consume()
	for msg := range c { // 接收消息
		object, e := strconv.Unquote(string(msg.Body)) // 消息的正文是接口服务发送过来的需要定位的对象名字。
		if e != nil {
			panic(e)
		}
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) { // 定位对象
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS")) // 向消息的发送方那个返回本服务节点的监听地址，表示该对象存在于本服务节点上。
		}
	}
}
