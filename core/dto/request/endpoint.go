/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2020/6/21 22:12
 */

package request

type CreateEndpointRequest struct {
	Address    string
	Weight     int
	UpstreamId string
}

type UpdateEndpointRequest struct {
	Id         string
	Address    string
	Weight     int
	UpstreamId string
}

type DeleteEndpointRequest struct {
	IdRequest
	UpstreamId string
}
