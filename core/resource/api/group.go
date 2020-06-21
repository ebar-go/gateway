/**
 * @Author: Hongker
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2020/6/14 13:28
 */

package api

import (
	"fmt"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/component/log"
)

// Group
type Group struct {
	// 这个radix tree是最重要的结构
	// 按照method将所有的方法分开, 然后每个method下面都是一个radix tree
	trees map[string]*node

	// 当/foo/没有匹配到的时候, 是否允许重定向到/foo路径
	RedirectTrailingSlash bool

	// 是否允许修正路径
	RedirectFixedPath bool

	// 如果当前无法匹配, 那么检查是否有其他方法能match当前的路由
	HandleMethodNotAllowed bool

	// 接口数量
	count int
}

func NewGroup() *Group {
	return &Group{
		trees:                  make(map[string]*node),
		RedirectTrailingSlash:  false,
		RedirectFixedPath:      false,
		HandleMethodNotAllowed: false,
		count:                  0,
	}
}

// Get return api with method,path
func (group *Group) Get(method, path string) *Api {
	api := new(Api)
	if root := group.trees[method]; root != nil {
		api.Method = method
		api.Path = path
		api.Key = root.getValue(path).(string)
	}
	return api
}

// Count return total count of api items
func (group *Group) Count() int {
	return group.count
}

// Add 添加
func (group *Group) Add(method, path, key string) error {
	if err := group.handle(method, path, key); err != nil {
		return fmt.Errorf("failed to add:%s", err.Error())
	}
	group.count++
	return nil
}

// Delete 删除
func (group *Group) Delete(method, path string) error {
	if err := group.handle(method, path, nil); err != nil {
		return fmt.Errorf("failed to delete:%s", err.Error())
	}

	group.count--
	return nil
}

// Handle registers a new request handle with the given path and method.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func (group *Group) handle(method, path string, handle Handle) error {
	if path[0] != '/' {
		return fmt.Errorf("path must begin with '/' in path '%s'", path)
	}

	if group.trees == nil {
		group.trees = make(map[string]*node)
	}

	root := group.trees[method]
	if root == nil {
		root = new(node)
		group.trees[method] = root
	}

	// 处理路由添加异常
	defer func() {
		if err := recover(); err != nil {
			app.Logger().Error("add router error", log.Context{
				"err": err,
			})
		}
	}()

	return root.addRoute(path, handle)
}
