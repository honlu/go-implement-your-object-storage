package heartbeat

import (
	"lib/rabbitmq"
	"os"
	"strconv"
	"sync"
	"time"
)

var dataServers = make(map[string]time.Time) // 在整个包中可见，用于缓存所有的数据服务节点。
var mutex sync.Mutex

/*
创建一个rabbitmq.RabbitMQ结构体，即消息队列来绑定apiServers exchange,
通过go channel监听每一个来自数据服务节点的心跳消息，
将该消息的正文内容，即数据服务节点的监听地址作为map的键，收到消息的时间作为值存入map。
*/
func ListenHeartbeat() {
	// fmt.Println(os.Getenv("RABBITMQ_SERVER")) // 测试
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

/*
每隔5s扫描一遍map，清除其中超过10s没收到心跳消息的数据服务节点。
*/
func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

/*
遍历map，返回当前所有的数据服务节点。
*/
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	return ds
}
