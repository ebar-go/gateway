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
	Id        int `gorm:"primary_key;column:id;type:int(10) auto_increment" json:"id"`
	CreatedAt int64 `gorm:"autoUpdateTime;not null;type:int(10);default:0" json:"created_at"`
	UpdateAt  int64 `gorm:"autoCreateTime;not null;type:int(10);default:0" json:"update_at"`
}
