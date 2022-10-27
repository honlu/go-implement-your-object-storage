package locate

import (
	"lib/rabbitmq"

	"os"
	"strconv"
	"time"
)

/*
接收一个string类型的参数name,即需要定位的对象的名字。
首先创建一个新的临时消息队列，并向dataServers exchanges群发这个对象名字的定位信息，
然后用goroutine启动一个匿名函数，用于在1s后关闭这个临时消息队列。
如果1s内有来自数据服务节点的消息，则返回消息的正文内容，即该数据服务节点的监听地址。
*/
func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

// 用于检测Locate结果是否为空字符串，来判断对象是否存在
func Exist(name string) bool {
	return Locate(name) != ""
}
