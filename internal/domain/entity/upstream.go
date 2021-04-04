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
	Name        string `gorm:"not null;type:varchar(50);default:''" json:"name"`
	Router      string `gorm:"not null;type:varchar(30);default:''" json:"router"`
	Status      int `gorm:"not null;type:tinyint(1);default:1" json:"status"`
	Description string `gorm:"not null;type:varchar(255);default:''" json:"description"`
}

func (UpstreamEntity) TableName() string {
	return TableUpstream
}