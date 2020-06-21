/**
 * @Author: Hongker
 * @Description:
 * @File:  api
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:43
 */

package entity

import (
	"fmt"
	"github.com/ebar-go/ego/utils/json"
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

func (e ApiEntity) PrimaryKey() string {
	// upstream/{uniqueUpstreamId}/endpoint/{uniqueApiId}
	return fmt.Sprintf("%s/%s/%s/%s", TableUpstream, e.UpstreamId, e.TableName(), e.Id)
}

func (e ApiEntity) Json() string {
	s, _ := json.Encode(e)
	return s
}
