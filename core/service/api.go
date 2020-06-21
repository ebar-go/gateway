/**
 * @Author: Hongker
 * @Description:
 * @File:  api
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

type apiService struct {
}

func Api() *apiService {
	return &apiService{}
}

// List 获取api列表
func (service *apiService) List(upstreamId string) ([]entity.ApiEntity, error) {
	upstream := &entity.UpstreamEntity{}
	upstream.Id = upstreamId

	items, err := dao.Api(app.Etcd()).List(upstream)
	if err != nil {
		return nil, errors.New(enum.DataQueryFailed, err.Error())
	}
	return items, nil

}

//
func (service *apiService) Create(req request.CreateApiRequest) error {
	item := &entity.ApiEntity{
		Method:     req.Method,
		Path:       req.Path,
		Key:        req.Key,
		UpstreamId: req.UpstreamId,
	}

	item.Id = strings.UUID()
	item.CreatedAt = date.GetTimeStamp()
	item.UpdateAt = date.GetTimeStamp()

	if err := dao.Api(app.Etcd()).Create(item); err != nil {
		return errors.New(enum.DataSaveFailed, err.Error())
	}
	return nil
}

func (service *apiService) Update(req request.UpdateApiRequest) error {
	upstream := &entity.UpstreamEntity{}
	upstream.Id = req.UpstreamId

	apiDao := dao.Api(app.Etcd())
	item, err := apiDao.Get(upstream, req.Id)
	if err != nil {
		return errors.New(enum.DataQueryFailed, err.Error())
	}

	item.Method = req.Method
	item.Path = req.Path
	item.Key = req.Key
	item.UpdateAt = date.GetTimeStamp()

	if err := dao.Api(app.Etcd()).Update(item); err != nil {
		return errors.New(enum.DataSaveFailed, err.Error())
	}
	return nil
}

// Delete 删除
func (service *apiService) Delete(upstreamId string, apiId string) error {
	upstream := &entity.UpstreamEntity{}
	upstream.Id = upstreamId

	apiDao := dao.Api(app.Etcd())
	item, err := apiDao.Get(upstream, apiId)
	if err != nil {
		return errors.New(enum.DataQueryFailed, err.Error())
	}

	if err := dao.Api(app.Etcd()).Delete(item); err != nil {
		return errors.New(enum.DataDeleteFailed, err.Error())
	}

	return nil
}
