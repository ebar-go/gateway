/**
 * @Author: Hongker
 * @Description:
 * @File:  interfaces
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:23
 */

package service

import "net/http"

type DispatcherService interface {
	DispatchRequest(router, path string, request *http.Request) (string, error)
}
