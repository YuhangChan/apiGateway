package main

import (
	"context"
	demo "github.com/SchrodingerwithCat/apiGateway/rpc/teacher_service/kitex_gen/demo"
)

// TeacherServiceImpl implements the last service interface defined in the IDL.
type TeacherServiceImpl struct{}

// TeacherRegister implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) TeacherRegister(ctx context.Context, teacher *demo.Teacher) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	return
}

// TeacherQuery implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) TeacherQuery(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...
	return
}
