/**
 * @Author: Hongker
 * @Description:
 * @File:  api
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:43
 */

package entity

import (
)

type ApiEntity struct {
	BaseEntity
	Method     string
	Path       string
	Key        string
	UpstreamId string
}

func (ApiEntity) TableName() string {
	return TableApi
}
