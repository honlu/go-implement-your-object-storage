package objects

import "net/http"

/*
w: 用于写入http的响应
r: 代表当前处理的HTTP的请求
*/
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method // Mehtod成员变量记录该HTTP请求的方法
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
