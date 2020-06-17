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
	base.PUT("upstream", nil)
	base.DELETE("upstream", handler.DeleteUpstreamHandler)
}
