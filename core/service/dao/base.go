/**
 * @Author: Hongker
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2020/6/17 23:11
 */

package dao

import (
	"context"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/component/etcd"
	"github.com/ebar-go/gateway/core/service/entity"
)

type BaseDao struct {
	conn *etcd.Client
}

func (dao *BaseDao) Create(item entity.Entity) error {
	_, err := app.Etcd().Api().Put(context.Background(), item.PrimaryKey(), item.Json())
	return err
}

func (dao *BaseDao) Update(item entity.Entity) error {
	_, err := app.Etcd().Api().Put(context.Background(), item.PrimaryKey(), item.Json())
	return err
}

func (dao *BaseDao) Delete(item entity.Entity) error {
	_, err := app.Etcd().Api().Delete(context.Background(), item.PrimaryKey())
	return err
}
