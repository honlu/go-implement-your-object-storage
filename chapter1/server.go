package main

import (
	"log"
	"net/http"
	"os"

	"chapter1/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)                    // 注册http处理函数
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil)) // 启动监听端口，一般函数永远不会返回。
}
