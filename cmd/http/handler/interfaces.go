/**
 * @Author: Hongker
 * @Description:
 * @File:  interfaces
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:15
 */

package handler

import "github.com/gin-gonic/gin"

type DispatcherHandler interface {
	Dispatch(ctx *gin.Context)
}
