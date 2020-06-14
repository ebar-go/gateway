/**
 * @Author: Hongker
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2020/6/14 11:03
 */

package handler

import (
	"github.com/ebar-go/ego/http/response"
	"github.com/gin-gonic/gin"
)

// UserLoginHandler 登录
func UserLoginHandler(ctx *gin.Context) {

}

// UserRegisterHandler 注册
func UserRegisterHandler(ctx *gin.Context) {
	response.WrapContext(ctx).Success(nil)
}
