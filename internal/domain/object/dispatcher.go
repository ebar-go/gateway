/**
 * @Author: Hongker
 * @Description:
 * @File:  dispatcher
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:37
 */

package object

import (
	"github.com/ebar-go/ego/component/curl"
	"net/http"
)

type Dispatcher interface {
	SendRequest(url string, req *http.Request) (string, error)
}

func NewDispatcher() Dispatcher {
	return &dispatcherImpl{}
}

type dispatcherImpl struct {

}

func (impl dispatcherImpl) buildRequest(url string, req *http.Request) (*http.Request, error) {
	return http.NewRequest(req.Method, url, req.Body)
}

func (impl dispatcherImpl) SendRequest(url string, req *http.Request) (string, error) {
	copyHttpRequest, err := impl.buildRequest(url, req)
	if err != nil {
		return "", err
	}

	response, err := curl.Send(copyHttpRequest)
	if err != nil {
		return "", err
	}
	return response.String(), nil
}

