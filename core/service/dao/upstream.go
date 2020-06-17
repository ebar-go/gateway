/**
 * @Author: Hongker
 * @Description:
 * @File:  upstream
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:47
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

type upstreamDao struct {
	BaseDao
}

func Upstream(conn *etcd.Client) *upstreamDao  {
	return &upstreamDao{
		BaseDao{conn:conn},
	}
}

func (dao upstreamDao) List() ([]entity.UpstreamEntity, error) {
	resp, err := dao.conn.Api().Get(context.Background(), entity.TableUpstream, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var items []entity.UpstreamEntity
	if resp == nil || resp.Kvs == nil {
		return nil, fmt.Errorf("Data not found")
	}

	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			var item entity.UpstreamEntity
			if err := json.Decode(v, &item); err != nil {
				continue
			}
			items = append(items, item)
		}
	}

	return items, nil
}

func (dao upstreamDao) Get(id string) (*entity.UpstreamEntity, error) {
	item := &entity.UpstreamEntity{}
	item.Id = id
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
