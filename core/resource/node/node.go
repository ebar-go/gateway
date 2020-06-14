package node

import (
	"fmt"
	"github.com/ebar-go/ego"
	"github.com/ebar-go/ego/component/trace"
	"github.com/ebar-go/gateway/core/resource/api"
	"github.com/ebar-go/gateway/core/resource/endpoint"
	"github.com/ebar-go/gateway/plugin/balance"
	"io/ioutil"
	"net/http"
	"strings"
)


// Node define properties of node
type Node struct {
	// unique id
	ID string

	// service name
	Router string

	// node status(0:offline,1:online)
	Status int

	// api
	ApiGroup *api.Group
	
	endpoints []endpoint.Endpoint

	// 权重负载均衡
	wb *balance.WeightBalance
}


// String
func (n *Node) String() string {
	return fmt.Sprintf("ID:%s, Router:%s, Status:%d, PodNumber:%d, ApiNumber:%d",
		n.ID, n.Router, n.Status, n.GetEndpointCount(),  n.ApiGroup.Count())
}

func (n *Node) getWeights() []int {
	var result []int
	for _, e := range n.endpoints {
		result = append(result, e.Weight())
	}

	return result
}

func (n *Node) GetEndpointCount() int {
	return len(n.endpoints)
}

func (n *Node) reloadWeight() {
	if n.wb == nil {
		n.wb = new(balance.WeightBalance)
	}
	n.wb.Reload(n.getWeights())
}

func (n *Node) AddEndpoint(endpoint endpoint.Endpoint) {
	n.endpoints = append(n.endpoints, endpoint)
	n.reloadWeight()
}

func (n *Node) DeleteEndpoint(id string) {
	for index, item := range n.endpoints {
		if item.Id() == id {
			n.endpoints = append(n.endpoints[:index], n.endpoints[index+1:]...)
		}
	}
	n.reloadWeight()
}


func (n *Node) getBackupIndex(lastIndex int) int {
	lastEndpoint := n.endpoints[lastIndex]
	n.DeleteEndpoint(lastEndpoint.Id())
	n.wb.Reload(n.getWeights())
	backupIndex := n.wb.RandomIndex()
	n.AddEndpoint(lastEndpoint)
	return backupIndex
}


func (n *Node) SendRequest(method, path string, req *http.Request) (string, error) {
	index := n.wb.RandomIndex()
	if index == balance.InvalidIndex {
		return "", fmt.Errorf("endpoint not found")
	}

	// 首次发送请求
	response, err := n.sendRequest(n.endpoints[index], method, path, req)
	if err == nil {
		return response, err
	}

	// 备用节点
	backupIndex := n.getBackupIndex(index)
	if backupIndex == balance.InvalidIndex {
		// not found backup endpoint
		return "", err
	}
	// 发起请求
	response, err = n.sendRequest(n.endpoints[backupIndex], method, path, req)
	if err != nil {
		return "", err
	}
	return response, nil
}

// SendRequest 执行请求
func (n *Node) sendRequest(c endpoint.Endpoint, method, path string, req *http.Request) (string, error) {
	// 构造http请求
	var request *http.Request
	var requestErr error
	var url = c.GetCompleteUrl(path)

	// 附加url参数
	if req.URL.RawQuery != "" {
		url = url + "?" + req.URL.RawQuery
	}
	switch method {
	case http.MethodPost, http.MethodPut:
		requestBody, _ := ioutil.ReadAll(req.Body)
		request, requestErr = http.NewRequest(req.Method, url, strings.NewReader(string(requestBody)))

	case http.MethodGet, http.MethodDelete:
		request, requestErr = http.NewRequest(req.Method, url, nil)

	default:
		requestErr = fmt.Errorf("method:%s is not support", req.Method)
	}

	if requestErr != nil {
		return "", requestErr
	}

	//设置请求类型
	request.Header.Set("Content-Type", req.Header.Get("Content-Type"))

	// 增加唯一TraceID
	traceId := req.Header.Get("Gateway-Trace")
	if traceId == "" {
		traceId = trace.Get()
	}
	request.Header.Add("Gateway-Trace", traceId)

	return ego.Curl(request)
}


