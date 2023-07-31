package idl

import (
	"github.com/cloudwego/kitex/pkg/generic"
	"os"
	"sync"
	"time"
)

var idlCache sync.Map
var cacheTime = make(map[*generic.ThriftContentProvider]time.Time)

func GetResolvedIdl(serviceName string) (*generic.ThriftContentProvider, error) {
	// 动态解析
	path := GetIdlPath(serviceName)
	// resolve the idl file's path
	content := GetIdlContent(serviceName)
	includes := map[string]string{
		path: content,
	}

	p, err := generic.NewThriftContentProvider(content, includes)
	if err != nil {
		panic(err)
	} else {
		go func() {
			err = p.UpdateIDL(content, includes)
			if err != nil {
				panic(err)
			}
			time.Sleep(30 * time.Second)
		}()
	}

	return p, err
}

// GetCacheIdl 缓存解析后的IDL内容
func GetCacheIdl(serviceName string) (*generic.ThriftContentProvider, error) {
	if idl, ok := idlCache.Load(serviceName); ok {
		oldIdl := idl.(*generic.ThriftContentProvider)
		cacheTimeout := 30 * time.Minute

		if time.Since(cacheTime[oldIdl]) < cacheTimeout {
			// 缓存未过期，直接返回
			return oldIdl, nil
		}
		// 缓存已过期，删除旧缓存
		idlCache.Delete(serviceName)
	}

	// 缓存不存在，创建新缓存
	idl, err := GetResolvedIdl(serviceName)
	if err != nil {
		return nil, err
	}
	cacheTime[idl] = time.Now() // 记录缓存时间
	idlCache.Store(serviceName, idl)

	return idl, nil
}

func GetIdlPath(serviceName string) string {
	service := GetService(serviceName)
	return service.ServiceIdlName
}

func GetIdlContent(serviceName string) string {
	service := GetService(serviceName)
	ioMutex.Lock()
	defer ioMutex.Unlock()

	// 读取文件
	file, err := os.Open(service.ServiceIdlName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := os.ReadFile(service.ServiceIdlName)
	if err != nil {
		panic(err)
	}
	return string(content)

}
