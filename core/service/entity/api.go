/**
 * @Author: Hongker
 * @Description:
 * @File:  api
 * @Version: 1.0.0
 * @Date: 2020/6/17 21:43
 */

package entity

type ApiEntity struct {
	Id string
	Method string
	Path string
	Key string
	UpstreamId int
}



func (ApiEntity) TableName() string {
	return "api"
}