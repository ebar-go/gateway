/**
 * @Author: Hongker
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:12
 */

package route

import (
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/gateway/cmd/http/handler"
	"github.com/gin-gonic/gin"
)

func Loader(router *gin.Engine, dispatcherHandler handler.DispatcherHandler)  {
	router.GET("/", func(ctx *gin.Context) {
		response.WrapContext(ctx).Success(nil)
	})
	router.Any("/api/:router/*api", dispatcherHandler.Dispatch)
}
