package objects

import (
	"log"
	"net/http"
	"strings"
)

/* 
put:
首先从url中获取对象名，然后将r.Body和object作为参数调用storeObject。
*/
func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, e := storeObject(r.Body, object)
	if e != nil {
		log.Println(e)
	}
	w.WriteHeader(c)
}
