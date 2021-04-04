/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:44
 */

package entity

import(
	"fmt"
)

type EndpointEntity struct {
	BaseEntity
	Address    string `gorm:"not null;type:varchar(50);default:''" json:"address"`
	Weight     int `gorm:"not null;type:int(10);default:0" json:"weight"`
	UpstreamId string `gorm:"index;not null;type:int(10);default:0" json:"upstream_id"`
}

func (EndpointEntity) TableName() string {
	return TableEndpoint
}


// GetCompleteUrl
func (entity EndpointEntity) GetCompleteUrl(path string) string {
	return fmt.Sprintf("http://%s%s", entity.Address, path)
}
