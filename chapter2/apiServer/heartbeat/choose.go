package heartbeat

import (
	"math/rand"
)

func ChooseRandomDataServer() string {
	ds := GetDataServers() // 这个函数怎么引用呢？
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
