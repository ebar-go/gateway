/**
 * @Author: Hongker
 * @Description:
 * @File:  api
 * @Version: 1.0.0
 * @Date: 2020/6/14 12:31
 */

package service

import (
	"github.com/ebar-go/gateway/core/dispatcher"
	"github.com/ebar-go/gateway/core/resource/api"
	"github.com/ebar-go/gateway/core/resource/endpoint"
	"github.com/ebar-go/gateway/core/resource/node"
	"net/http"
)

var httpDispatcher  = dispatcher.NewHttpDispatcher()
type dispatcherService struct {

}

func init()  {

	apiGroup := api.NewGroup()
	_ = apiGroup.Add(http.MethodGet, "/user", "v1.user.list")

	nodeInstance := &node.Node{
		ID:       "1",
		Router:   "user",
		Status:   node.Online,
		ApiGroup: apiGroup,
	}

	nodeInstance.AddEndpoint(endpoint.New( "127.0.0.1:9001", 10))
	nodeInstance.AddEndpoint(endpoint.New( "127.0.0.1:9002", 20))

	httpDispatcher.NodeGroup().Add(nodeInstance)

}

func Dispatcher() *dispatcherService {
	return &dispatcherService{}
}

func (service dispatcherService) Dispatch(router, path string, request *http.Request) (string, error) {
	return httpDispatcher.Dispatch(router, path, request)
}
