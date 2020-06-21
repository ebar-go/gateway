/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:40
 */

package service

import (
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/utils/date"
	"github.com/ebar-go/ego/utils/strings"
	"github.com/ebar-go/gateway/core/dto/request"
	"github.com/ebar-go/gateway/core/enum"
	"github.com/ebar-go/gateway/core/service/dao"
	"github.com/ebar-go/gateway/core/service/entity"
)

type endpointService struct {
}

func Endpoint() *endpointService {
	return &endpointService{}
}

// List 获取endpoint列表
func (service *endpointService) List(upstreamId string) ([]entity.EndpointEntity, error) {
	upstream := &entity.UpstreamEntity{}
	upstream.Id = upstreamId

	items, err := dao.Endpoint(app.Etcd()).List(upstream)
	if err != nil {
		return nil, errors.New(enum.DataQueryFailed, err.Error())
	}
	return items, nil

}

//
func (service *endpointService) Create(req request.CreateEndpointRequest) error {
	item := &entity.EndpointEntity{
		Address:    req.Address,
		Weight:     req.Weight,
		UpstreamId: req.UpstreamId,
	}

	item.Id = strings.UUID()
	item.CreatedAt = date.GetTimeStamp()
	item.UpdateAt = date.GetTimeStamp()

	if err := dao.Endpoint(app.Etcd()).Create(item); err != nil {
		return errors.New(enum.DataSaveFailed, err.Error())
	}
	return nil
}

func (service *endpointService) Update(req request.UpdateEndpointRequest) error {
	upstream := &entity.UpstreamEntity{}
	upstream.Id = req.UpstreamId

	endpointDao := dao.Endpoint(app.Etcd())
	item, err := endpointDao.Get(upstream, req.Id)
	if err != nil {
		return errors.New(enum.DataQueryFailed, err.Error())
	}

	item.Address = req.Address
	item.Weight = req.Weight
	item.UpdateAt = date.GetTimeStamp()

	if err := dao.Endpoint(app.Etcd()).Update(item); err != nil {
		return errors.New(enum.DataSaveFailed, err.Error())
	}
	return nil
}

// Delete 删除
func (service *endpointService) Delete(upstreamId string, endpointId string) error {
	upstream := &entity.UpstreamEntity{}
	upstream.Id = upstreamId

	endpointDao := dao.Endpoint(app.Etcd())
	item, err := endpointDao.Get(upstream, endpointId)
	if err != nil {
		return errors.New(enum.DataQueryFailed, err.Error())
	}

	if err := dao.Endpoint(app.Etcd()).Delete(item); err != nil {
		return errors.New(enum.DataDeleteFailed, err.Error())
	}

	return nil
}
