// Code generated by hertz generator.

package demo

import (
	"context"
	"encoding/json"
	"fmt"
	client_provider "github.com/SchrodingerwithCat/apiGateway/http/biz/clientprovider"
	"github.com/cloudwego/kitex/client/genericclient"

	demo "github.com/SchrodingerwithCat/apiGateway/http/biz/model/demo"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// TeacherRegister .
// @router /teacher/add-teacher-info [POST]
func TeacherRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.Student
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	jsonReq, err := json.Marshal(req)

	if err != nil {
		fmt.Println("error:", err)
	}

	cli := client_provider.GetTeacherGenericClient(&ctx, c).(genericclient.Client)
	resp, err := cli.GenericCall(ctx, "TeacherRegister", string(jsonReq))
	c.JSON(consts.StatusOK, resp)
}

// TeacherQuery .
// @router /teacher/query [GET]
func TeacherQuery(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.QueryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	jsonReq, err := json.Marshal(req)

	if err != nil {
		fmt.Println("error:", err)
	}

	cli := client_provider.GetTeacherGenericClient(&ctx, c).(genericclient.Client)
	resp, err := cli.GenericCall(ctx, "TeacherQuery", string(jsonReq))
	c.JSON(consts.StatusOK, resp)
}
