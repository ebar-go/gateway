/**
 * @Author: Hongker
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2020/6/14 13:03
 */

package upstream

import (
	"fmt"
)

type Group struct {
	items map[string]*Upstream
}

func NewGroup()  *Group{
	return &Group{items: map[string]*Upstream{}}
}

func (group *Group) Check(project Upstream) error {
	if group.Has(project.ID) {
		return fmt.Errorf("upstream:%s is exists", project.ID)
	}

	if group.HasRouter(project.Router) {
		return fmt.Errorf("upstream router:%s has been used", project.Router)
	}
	return nil
}

// Add add the upstream to group
func (group *Group) Add(upm *Upstream) {
	group.items[upm.ID] = upm
}

// Has check upstream exist
func (group *Group) Has(id string) bool {
	_, ok := group.items[id]
	return ok
}

// HasRouter if upstream router has been used, return true
func (group *Group) HasRouter(router string) bool {
	for _, item := range group.items {
		if item.Router == router {
			return true
		}
	}

	return false
}

// FindByRouter
func (group *Group) FindByRouter(router string) *Upstream {
	var project *Upstream
	for _, item := range group.items {
		if item.Router == router {
			project = item
			break
		}
	}

	return project
}

// Get return upstream if exist
func (group *Group) Get(id string) *Upstream {
	return group.items[id]
}

// Delete
func (group *Group) Delete(id string) {
	delete(group.items, id)
}
