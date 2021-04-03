/**
 * @Author: Hongker
 * @Description:
 * @File:  upstream
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:42
 */

package entity

import (
)

type UpstreamEntity struct {
	BaseEntity
	Name        string `json:"name"`
	Router      string `json:"router"`
	Status      int `json:"status"`
	Description string `json:"description"`
}

func (UpstreamEntity) TableName() string {
	return TableUpstream
}