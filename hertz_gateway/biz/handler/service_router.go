package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"hertz.demo/biz/clientprovider"
)

func CallServiceMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interface{}
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	serviceName := c.Param("service")
	methodName := c.Param("method")
	fmt.Println("Calling the service: " + serviceName + "'s method: " + methodName)

	// 将请求req参数转换为 JSON形式的数据
	jsonReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("jsonReq:", string(jsonReq))

	// 通过本地再次生成的形式添加客户端，没有上传到 GitHub 获取对应service的客户端
	cli, err := clientprovider.GetCacheClient(serviceName)

	// 泛化调用
	resp, err := cli.GenericCall(ctx, methodName, string(jsonReq))

	c.JSON(consts.StatusOK, resp)
}
