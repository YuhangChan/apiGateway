package main

import (
	"github.com/Qi118/cloudwego/kitex_gen/teacher/teacherservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)


func initEtcd(handler *TeacherServiceImpl, addr *net.TCPAddr) server.Server {

	r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})
	if err != nil {
		log.Println("fail to init etcd registry" + err.Error())
	}

	ebi := &rpcinfo.EndpointBasicInfo{
		ServiceName: "teacher-server",
		Tags:        map[string]string{"Cluster": "teacher"},
	}

	svr := teacherservice.NewServer(handler, server.WithRegistry(r), server.WithServiceAddr(addr), server.WithServerBasicInfo(ebi))

	return svr
}

func main() {
    addr, err := net.ResolveTCPAddr("tcp", :9997)
    if err != nil {
	    log.Println(err.Error())
    }

	r, err := etcd.etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
	    log.Println(err.Error())
    }

	


	handler := new(TeacherServiceImpl)

	svr := initEtcd(handler, addr)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
