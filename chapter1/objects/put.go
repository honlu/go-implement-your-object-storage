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
func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2]) // 首先获取经过转义后的路径部分，然后创建一个同名文件f
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError) // 创建失败，返回500
		return
	}
	defer f.Close()    // 延迟关闭文件
	io.Copy(f, r.Body) // 将r.Body写入文件f
}
