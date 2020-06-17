/**
 * @Author: Hongker
 * @Description: http请求转发器
 * @File:  dispatcher
 * @Version: 1.0.0
 * @Date: 2020/6/14 11:26
 */

package dispatcher

import (
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/gateway/core/enum"
	"github.com/ebar-go/gateway/core/resource/upstream"
	"net/http"
)

type Dispatcher interface {
	Dispatch(router, path string, request *http.Request) (string, error)
}

type HttpDispatcher struct {
	ng *upstream.Group
}

func NewHttpDispatcher() *HttpDispatcher {
	return &HttpDispatcher{ng: upstream.NewGroup()}
}

func (dispatcher *HttpDispatcher) UpstreamGroup() *upstream.Group {
	return dispatcher.ng
}

// Dispatch 转发
func (dispatcher *HttpDispatcher) Dispatch(router, path string, request *http.Request) (string, error) {
	n := dispatcher.ng.FindByRouter(router)
	if n == nil {
		return "", errors.New(enum.DataNotFound, "upstream not found")
	}

	if n.Status == upstream.Offline {
		return "", errors.New(enum.UpstreamWasOffline, "upstream is offline")
	}

	api := n.ApiGroup.Get(request.Method, path)
	if api == nil {
		return "", errors.New(enum.DataNotFound, "api resource not found")
	}

	return n.SendRequest(api.Method, api.Path, request)
}
