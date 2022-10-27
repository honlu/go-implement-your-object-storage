package main

import (
	"log"
	"net/http"
	"os"

	"apiServer/heartbeat" // 注意写法：本项目包下
	"apiServer/locate"
	"apiServer/objects"
)

/*
接口服务
heartbeat包：用于接收数据服务节点的心跳信息。
locate包：用于向数据服务节点群发定位信息并接收反馈消息。
objects包：负责将对象请求转发给数据服务。
*/
func main() {
	go heartbeat.ListenHeartbeat() // 开启一个协程
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
