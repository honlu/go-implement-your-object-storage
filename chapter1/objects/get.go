package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
w: 用于写入http的响应
r: 代表当前处理的HTTP的请求
*/
func get(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2]) // 首先解析路径，然后打开文件
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound) // 打开失败，返回404
		return
	}
	defer f.Close()
	io.Copy(w, f) // 打开成功，将f的内容写入w
}
