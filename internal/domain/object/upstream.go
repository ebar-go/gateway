/**
 * @Author: Hongker
 * @Description:
 * @File:  upstream
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:37
 */

package object

type Upstream interface {
	GetEndpoint() (Endpoint, error)
}

type upstreamImpl struct {

}
