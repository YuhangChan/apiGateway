package main

import (
	demo "github.com/SchrodingerwithCat/apiGateway/rpc/teacher_service/kitex_gen/demo/teacherservice"
	"log"
)

func main() {
	svr := demo.NewServer(new(TeacherServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
