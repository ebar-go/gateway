/**
 * @Author: Hongker
 * @Description:
 * @File:  upstream
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:42
 */

package entity

import (
	"fmt"
	"github.com/ebar-go/ego/utils/json"
)

type UpstreamEntity struct {
	BaseEntity
	Name        string
	Router      string
	Status      int
	Description string
}

func (UpstreamEntity) TableName() string {
	return TableUpstream
}

func (e UpstreamEntity) PrimaryKey() string {
	return fmt.Sprintf("%s/%s", e.TableName(), e.Id)
}

func (e UpstreamEntity) Json() string {
	s, _ := json.Encode(e)
	return s
}

// GetEndpointPrefix
func (e UpstreamEntity) GetEndpointPrefix() string {
	return fmt.Sprintf("%s/%s/%s", e.TableName(), e.Id, TableEndpoint)
}

// GetApiPrefix
func (e UpstreamEntity) GetApiPrefix() string {
	return fmt.Sprintf("%s/%s/%s", e.TableName(), e.Id, TableApi)
}
