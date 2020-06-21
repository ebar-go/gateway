/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2020/6/14 13:17
 */

package endpoint

import (
	"fmt"
	"github.com/ebar-go/ego/utils/strings"
)

// Endpoint
type Endpoint struct {
	id string
	// http Address
	address string
	// weight
	weight int
}

func New(Address string, weight int) Endpoint {
	return Endpoint{id: strings.UUID(), address: Address, weight: weight}
}

// GetCompleteUrl
func (e Endpoint) GetCompleteUrl(path string) string {
	return fmt.Sprintf("http://%s%s", e.address, path)
}

// Id
func (e Endpoint) Id() string {
	return e.id
}

// Weight 权重
func (e Endpoint) Weight() int {
	return e.weight
}
