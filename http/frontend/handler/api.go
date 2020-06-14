/**
 * @Author: Hongker
 * @Description:
 * @File:  api
 * @Version: 1.0.0
 * @Date: 2020/6/14 11:03
 */

package handler

import (
	"github.com/ebar-go/ego/utils/secure"
	"github.com/ebar-go/gateway/core/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DispatchHandler 转发请求
func DispatchHandler(ctx *gin.Context) {
	// 1.获取服务名称与uri
	router := ctx.Param("router")
	path := ctx.Param("api")

	// 序列化
	result, err := service.Dispatcher().Dispatch(router, path, ctx.Request)
	secure.Panic(err)

	ctx.String(http.StatusOK, result)
}
