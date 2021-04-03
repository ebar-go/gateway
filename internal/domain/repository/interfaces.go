/**
 * @Author: Hongker
 * @Description:
 * @File:  interfaces
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:31
 */

package repository

import (
	"github.com/ebar-go/gateway/internal/domain/entity"
	"github.com/ebar-go/gateway/internal/domain/object"
)

type DispatcherRepo interface {
	GetHttpDispatcher() object.Dispatcher
}


type EndpointRepo interface {
	FindByUpstream(upstream *entity.UpstreamEntity) (*entity.EndpointEntity, error)
}

type UpstreamRepo interface {
	FindByRouter(router string) (*entity.UpstreamEntity, error)
}