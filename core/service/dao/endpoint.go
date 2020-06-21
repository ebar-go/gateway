/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2020/6/21 21:44
 */

package dao

import (
	"context"
	"fmt"
	"github.com/ebar-go/ego/component/etcd"
	"github.com/ebar-go/ego/utils/json"
	"github.com/ebar-go/gateway/core/service/entity"
	"go.etcd.io/etcd/clientv3"
)

type endpointDao struct {
	BaseDao
}

func Endpoint(conn *etcd.Client) *endpointDao {
	return &endpointDao{
		BaseDao{conn: conn},
	}
}

// List
func (dao endpointDao) List(upstream *entity.UpstreamEntity) ([]entity.EndpointEntity, error) {
	resp, err := dao.conn.Api().Get(context.Background(), upstream.GetEndpointPrefix(), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var items []entity.EndpointEntity
	if resp == nil || resp.Kvs == nil {
		return nil, fmt.Errorf("Data not found")
	}

	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			var item entity.EndpointEntity
			if err := json.Decode(v, &item); err != nil {
				continue
			}
			items = append(items, item)
		}
	}

	return items, nil
}

func (dao *endpointDao) Get(upstream *entity.UpstreamEntity, id string) (*entity.EndpointEntity, error) {
	item := &entity.EndpointEntity{}
	item.Id = id
	item.UpstreamId = upstream.Id
	resp, err := dao.conn.Api().Get(context.Background(), item.PrimaryKey())
	if err != nil {
		return nil, err
	}

	if resp == nil || resp.Kvs == nil {
		return nil, fmt.Errorf("Data not found")
	}

	if v := resp.Kvs[0].Value; v != nil {
		if err := json.Decode(v, item); err != nil {
			return nil, err
		}
	}
	return item, nil
}
