/**
 * @Author: Hongker
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2020/6/14 11:06
 */

package frontend

import (
	"github.com/ebar-go/gateway/http/frontend/handler"
	"github.com/gin-gonic/gin"
)

// LoadRoute 加载路由
func LoadRoute(router *gin.Engine) {
	base := router.Group("v1/frontend")

	base.Any("/api/:router/*api", handler.DispatchHandler)
}
