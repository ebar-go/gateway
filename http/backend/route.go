/**
 * @Author: Hongker
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2020/6/14 11:06
 */

package backend

import (
	"github.com/ebar-go/gateway/http/backend/handler"
	"github.com/gin-gonic/gin"
)

// LoadRoute 加载路由
func LoadRoute(router *gin.Engine) {
	base := router.Group("v1/backend")

	base.POST("user/login", handler.UserLoginHandler)
	base.POST("user/register", handler.UserRegisterHandler)

	base.GET("upstream", handler.ListUpstreamHandler)
	base.POST("upstream", handler.CreateUpstreamHandler)
	base.PUT("upstream", handler.UpdateUpstreamHandler)
	base.DELETE("upstream", handler.DeleteUpstreamHandler)

	base.GET("endpoint", handler.ListEndpointHandler)
	base.POST("endpoint", handler.CreateEndpointHandler)
	base.PUT("endpoint", handler.UpdateEndpointHandler)
	base.DELETE("endpoint", handler.DeleteEndpointHandler)

	base.GET("api", handler.ListApiHandler)
	base.POST("api", handler.CreateApiHandler)
	base.PUT("api", handler.UpdateApiHandler)
	base.DELETE("api", handler.DeleteApiHandler)

}
