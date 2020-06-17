/**
 * @Author: Hongker
 * @Description:
 * @File:  upstreamServer
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:37
 */

package service

import (
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/utils/date"
	"github.com/ebar-go/ego/utils/strings"
	"github.com/ebar-go/gateway/core/dto/request"
	"github.com/ebar-go/gateway/core/enum"
	"github.com/ebar-go/gateway/core/resource/upstream"
	"github.com/ebar-go/gateway/core/service/dao"
	"github.com/ebar-go/gateway/core/service/entity"
)

type upstreamService struct {

}

func Upstream() *upstreamService {
	return &upstreamService{}
}

func (service upstreamService) List() ([]entity.UpstreamEntity, error){
	return dao.Upstream(app.Etcd()).List()
}

func (service upstreamService) Create(req request.CreateUpstreamRequest) error  {
	item := &entity.UpstreamEntity{
		Name:        req.Name,
		Router:      req.Router,
		Status:      upstream.Offline,
		Description: req.Description,
	}
	item.Id = strings.UUID()
	item.CreatedAt = date.GetTimeStamp()
	item.UpdateAt = date.GetTimeStamp()

	if err := dao.Upstream(app.Etcd()).Create(item); err != nil {
		return errors.New(enum.DataSaveFailed, err.Error())
	}
	return nil
}

func (service upstreamService) Update() error {
	return nil
}

// Delete 删除
func (service upstreamService) Delete(id string) error  {
	upstreamDao := dao.Upstream(app.Etcd())
	item, err := upstreamDao.Get(id)
	if err != nil {
		return errors.New(enum.DataNotFound, err.Error())
	}

	if err := upstreamDao.Delete(item); err != nil {
		return errors.New(enum.DataDeleteFailed, err.Error())
	}
	return nil
}
