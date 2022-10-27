package main

import (
	"dataServer/heartbeat"
	"dataServer/locate"
	"dataServer/objects"

	"log"
	"net/http"
	"os"
)

/*
数据服务main函数：比单机版多了两个goroutine.goroutine是go并发执行的模型。
heartbeat包：心跳包，用于发送心跳信息。
locate包：定位包，用于接收定位消息、定位对象以及发送反馈消息。
objects包：对象的get,put操作。负责对象在本地磁盘上的存取。
*/
func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
