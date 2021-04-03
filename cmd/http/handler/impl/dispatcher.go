/**
 * @Author: Hongker
 * @Description:
 * @File:  dispatcher
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:16
 */
package impl

import (
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/gateway/cmd/http/handler"
	"github.com/ebar-go/gateway/internal/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// dispatcherHandlerImpl implement of handler.DispatcherHandler
type dispatcherHandlerImpl struct {
	service service.DispatcherService
}

func (impl dispatcherHandlerImpl) Dispatch(ctx *gin.Context) {
	// 1.获取服务名称与uri
	router := ctx.Param("router")
	path := ctx.Param("api")

	// 序列化
	result, err := impl.service.DispatchRequest(router, path, ctx.Request)
	if err != nil {
		response.WrapContext(ctx).Error(1001, err.Error())
		return
	}

	ctx.String(http.StatusOK, result)
}

func newDispatcherHandler() handler.DispatcherHandler {
	return new(dispatcherHandlerImpl)
}

