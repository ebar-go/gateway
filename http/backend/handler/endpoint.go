/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2020/6/21 22:26
 */

package handler

import (
	"fmt"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/ebar-go/gateway/core/dto/request"
	"github.com/ebar-go/gateway/core/enum"
	"github.com/ebar-go/gateway/core/service"
	"github.com/gin-gonic/gin"
)

func ListEndpointHandler(ctx *gin.Context) {
	upstreamId := ctx.Query("upstream_id")
	items, err := service.Endpoint().List(upstreamId)
	secure.Panic(err)
	response.WrapContext(ctx).Success(items)
}

func CreateEndpointHandler(ctx *gin.Context) {
	var req request.CreateEndpointRequest
	if err := ctx.ShouldBind(&req); err != nil {
		secure.Panic(errors.New(enum.InvalidParam, fmt.Sprintf("参数错误:%v", err)))
	}

	secure.Panic(service.Endpoint().Create(req))
	response.WrapContext(ctx).Success(nil)
}

func UpdateEndpointHandler(ctx *gin.Context) {
	var req request.UpdateEndpointRequest
	if err := ctx.ShouldBind(&req); err != nil {
		secure.Panic(errors.New(enum.InvalidParam, fmt.Sprintf("参数错误:%v", err)))
	}

	secure.Panic(service.Endpoint().Update(req))
	response.WrapContext(ctx).Success(nil)
}

func DeleteEndpointHandler(ctx *gin.Context) {
	var req request.DeleteEndpointRequest
	if err := ctx.ShouldBind(&req); err != nil {
		secure.Panic(errors.New(enum.InvalidParam, fmt.Sprintf("参数错误:%v", err)))
	}

	secure.Panic(service.Endpoint().Delete(req.UpstreamId, req.Id))
	response.WrapContext(ctx).Success(nil)
}
