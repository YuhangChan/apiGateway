// idl management platform - backend

package idl

import (
	"errors"
	"hertz.demo/biz/model/gateway"
	"sync"
)

var ServiceNameMap = make(map[string]gateway.Service)

var mapMutex = &sync.Mutex{}
var ioMutex = &sync.Mutex{}

func GetService(serviceName string) *gateway.Service {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	service, ok := ServiceNameMap[serviceName]
	if !ok {
		err := errors.New("service not found")
		panic(err)
	}
	return &service
}
