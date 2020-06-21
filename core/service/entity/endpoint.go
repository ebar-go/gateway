/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:44
 */

package entity

import (
	"fmt"
	"github.com/ebar-go/ego/utils/json"
)

type EndpointEntity struct {
	BaseEntity
	Address    string
	Weight     int
	UpstreamId string
}

func (EndpointEntity) TableName() string {
	return TableEndpoint
}

func (e EndpointEntity) PrimaryKey() string {
	// upstream/{uniqueUpstreamId}/endpoint/{uniqueEndpointId}
	return fmt.Sprintf("%s/%s/%s/%s", TableUpstream, e.UpstreamId, e.TableName(), e.Id)
}

func (e EndpointEntity) Json() string {
	s, _ := json.Encode(e)
	return s
}
