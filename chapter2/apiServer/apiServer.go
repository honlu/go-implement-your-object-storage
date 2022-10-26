package main

import (
	"log"
	"net/http"
	"os"

	"apiServer/heartbeat" // 注意写法：本项目包下
	"apiServer/locate"
	"apiServer/objects"
)

func main() {
	go heartbeat.ListenHeartbeat() // 开启一个协程
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
