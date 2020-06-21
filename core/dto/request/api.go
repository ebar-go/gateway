/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2020/6/21 22:12
 */

package request

type CreateApiRequest struct {
	Method     string
	Path       string
	Key        string
	UpstreamId string
}

type UpdateApiRequest struct {
	Id         string
	Method     string
	Path       string
	Key        string
	UpstreamId string
}

type DeleteApiRequest struct {
	IdRequest
	UpstreamId string
}
