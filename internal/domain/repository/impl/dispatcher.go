/**
 * @Author: Hongker
 * @Description:
 * @File:  dispatcher
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:32
 */

package impl

import (
	"github.com/ebar-go/gateway/internal/domain/object"
	"github.com/ebar-go/gateway/internal/domain/repository"
	"sync"
)

type dispatcherRepoImpl struct {
	// 使用线程池，减少GC
	pool *sync.Pool
}

func (impl dispatcherRepoImpl) GetHttpDispatcher() object.Dispatcher {
	return object.NewDispatcher()
}

func newDispatcherRepo() repository.DispatcherRepo {
	return &dispatcherRepoImpl{pool: new(sync.Pool)}
}
