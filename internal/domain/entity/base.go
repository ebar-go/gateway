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
	Id        int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdateAt  int64 `json:"update_at"`
}
