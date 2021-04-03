/**
 * @Author: Hongker
 * @Description:
 * @File:  dispatcher
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:29
 */

package impl

import (
	"fmt"
	"github.com/ebar-go/gateway/internal/domain/repository"
	"github.com/ebar-go/gateway/internal/domain/service"
	"net/http"
)

type dispatcherServiceImpl struct {
	dispatcherRepo repository.DispatcherRepo
	endpointRepo repository.EndpointRepo
	upstreamRepo repository.UpstreamRepo
}

func (impl dispatcherServiceImpl) DispatchRequest(router, path string, request *http.Request) (string, error) {
	upstream, err := impl.upstreamRepo.FindByRouter(router)
	if err != nil {
		return "", fmt.Errorf("查找上游服务失败")
	}

	endpoint, err := impl.endpointRepo.FindByUpstream(upstream)
	if err != nil {
		return "", fmt.Errorf("获取节点信息失败")
	}

	dispatcher := impl.dispatcherRepo.GetHttpDispatcher()
	return dispatcher.SendRequest(endpoint.GetCompleteUrl(path), request)


}

func newDispatcherService() service.DispatcherService {
	return &dispatcherServiceImpl{}
}
