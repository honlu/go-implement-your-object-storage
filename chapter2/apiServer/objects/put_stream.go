package objects

import (
	"fmt"
	"lib/objectstream"

	"apiServer/heartbeat"
)

/*
首先获取一个随机数据服务节点的地址server.
*/
func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewPutStream(server, object), nil
}
