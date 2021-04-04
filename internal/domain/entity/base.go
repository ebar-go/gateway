/**
 * @Author: Hongker
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2020/6/17 22:51
 */

package entity

const (
	TableUpstream = "upstream"
	TableEndpoint = "endpoint"
	TableApi      = "api"
)

type BaseEntity struct {
	Id        int `gorm:"primary_key;autoIncrement;column:id" json:"id"`
	CreatedAt int64 `gorm:"autoUpdateTime;not null" json:"created_at"`
	UpdateAt  int64 `gorm:"autoCreateTime" json:"update_at"`
}
