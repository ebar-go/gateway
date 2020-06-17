/**
 * @Author: Hongker
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2020/6/17 22:51
 */

package entity

const(
	TableUpstream = "upstream"
)

type Entity interface {
	TableName() string
	PrimaryKey() string
	Json() string
}

type BaseEntity struct {
	Id string
	CreatedAt int64
	UpdateAt int64
}