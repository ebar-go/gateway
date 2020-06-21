/**
 * @Author: Hongker
 * @Description:
 * @File:  upstream
 * @Version: 1.0.0
 * @Date: 2020/6/17 23:36
 */

package request

type CreateUpstreamRequest struct {
	Name        string
	Router      string
	Description string
}

type UpdateUpstreamRequest struct {
	Id          string
	Name        string
	Router      string
	Description string
}
