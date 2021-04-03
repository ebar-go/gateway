/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:43
 */

package impl

import (
	"fmt"
	"github.com/ebar-go/ego/component/mysql"
	"github.com/ebar-go/gateway/internal/domain/entity"
	"github.com/ebar-go/gateway/internal/domain/repository"
	"github.com/ebar-go/gateway/internal/plugin/balance"
)

type endpointRepoImpl struct {
	db mysql.Database
	balancer *balance.WeightBalance
}

func (impl endpointRepoImpl) FindByUpstream(upstream *entity.UpstreamEntity) (*entity.EndpointEntity, error) {
	items, err := impl.findAllByUpstreamId(upstream.Id)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("no endpoints")
	}

	item := impl.chooseOne(items)


	return &item, nil
}

func (impl endpointRepoImpl) chooseOne(items []entity.EndpointEntity) entity.EndpointEntity {
	if len(items) == 1 {
		return items[0]
	}

	weights := []int{}
	for _, item := range items {
		weights = append(weights, item.Weight)
	}
	// todo reload once
	impl.balancer.Reload(weights)
	index := impl.balancer.RandomIndex()

	return items[index]
}

func (impl endpointRepoImpl) findAllByUpstreamId(upstreamId int) ([]entity.EndpointEntity, error){
	items := make([]entity.EndpointEntity, 0)

	query := impl.db.GetInstance().Table(entity.TableEndpoint).Where("upstream_id = ?", upstreamId)
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func newEndpointRepo(db mysql.Database) repository.EndpointRepo {
	return &endpointRepoImpl{db: db, balancer: &balance.WeightBalance{}}
}
