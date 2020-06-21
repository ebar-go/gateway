/**
 * @Author: Hongker
 * @Description:
 * @File:  api
 * @Version: 1.0.0
 * @Date: 2020/6/14 12:31
 */

package service

import (
	"context"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/utils/json"
	"github.com/ebar-go/gateway/core/dispatcher"
	"github.com/ebar-go/gateway/core/resource/api"
	"github.com/ebar-go/gateway/core/resource/endpoint"
	"github.com/ebar-go/gateway/core/resource/upstream"
	"github.com/ebar-go/gateway/core/service/entity"
	"go.etcd.io/etcd/clientv3"
	"net/http"
)

var httpDispatcher = dispatcher.NewHttpDispatcher()

type dispatcherService struct {
}

func init() {

	apiGroup := api.NewGroup()
	_ = apiGroup.Add(http.MethodGet, "/user", "v1.user.list")

	nodeInstance := &upstream.Upstream{
		ID:       "1",
		Router:   "user",
		Status:   upstream.Online,
		ApiGroup: apiGroup,
	}

	nodeInstance.AddEndpoint(endpoint.New("127.0.0.1:9001", 10))
	nodeInstance.AddEndpoint(endpoint.New("127.0.0.1:9002", 20))

	httpDispatcher.UpstreamGroup().Add(nodeInstance)

}

func Dispatcher() *dispatcherService {
	return &dispatcherService{}
}

func (service dispatcherService) Dispatch(router, path string, request *http.Request) (string, error) {
	return httpDispatcher.Dispatch(router, path, request)
}

func (service dispatcherService) WatchUpstream() error {
	rch := app.Etcd().Instance().Watch(context.Background(), entity.TableUpstream, clientv3.WithPrefix())

	for wresp := range rch {
		for _, ev := range wresp.Events {
			id := string(ev.Kv.Key)
			switch ev.Type {
			case clientv3.EventTypePut:
				item := new(entity.UpstreamEntity)
				if err := json.Decode(ev.Kv.Value, item); err != nil {
					continue
				}

			case clientv3.EventTypeDelete:
				httpDispatcher.UpstreamGroup().Delete(id)
			}
		}
	}

	return nil
}

func (service dispatcherService) WatchEndpoint() error {
	return nil
}

func (service dispatcherService) WatchApi() error {
	return nil
}
