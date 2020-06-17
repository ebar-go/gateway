/**
 * @Author: Hongker
 * @Description:
 * @File:  upstream
 * @Version: 1.0.0
 * @Date: 2020/6/17 23:32
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

func ListUpstreamHandler(ctx *gin.Context)  {
	items, err := service.Upstream().List()
	secure.Panic(err)
	response.WrapContext(ctx).Success(items)
}

func CreateUpstreamHandler(ctx *gin.Context)  {
	var req request.CreateUpstreamRequest
	if err := ctx.ShouldBind(&req); err != nil {
		secure.Panic(errors.New(enum.InvalidParam, fmt.Sprintf("参数错误:%v", err)))
	}

	secure.Panic(service.Upstream().Create(req))
	response.WrapContext(ctx).Success(nil)
}


func DeleteUpstreamHandler(ctx *gin.Context)  {
	var req request.IdRequest
	if err := ctx.ShouldBind(&req); err != nil {
		secure.Panic(errors.New(enum.InvalidParam, fmt.Sprintf("参数错误:%v", err)))
	}

	secure.Panic(service.Upstream().Delete(req.Id))
	response.WrapContext(ctx).Success(nil)
}