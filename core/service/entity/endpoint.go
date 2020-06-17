/**
 * @Author: Hongker
 * @Description:
 * @File:  endpoint
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:44
 */

package entity

type EndpointEntity struct {
	Id string
	Address string
	Weight int
	UpstreamId int
}



func (EndpointEntity) TableName() string {
	return "endpoint"
}