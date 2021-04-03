/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:39
 */

package object

import (
	"fmt"
	"github.com/ebar-go/egu"
)

type Endpoint interface {
	GetCompleteUrl(path string) string
	Weight() int
	Id() string
}

type endpointImpl struct {
	id string
	// http Address
	address string
	// weight
	weight int
}


func New(Address string, weight int) Endpoint {
	return &endpointImpl{id:egu.UUID(), address: Address, weight: weight}
}

// GetCompleteUrl
func (impl endpointImpl) GetCompleteUrl(path string) string {
	return fmt.Sprintf("http://%s%s", impl.address, path)
}

// Id
func (impl endpointImpl) Id() string {
	return impl.id
}

// Weight 权重
func (impl endpointImpl) Weight() int {
	return impl.weight
}
