/**
 * @Author: Hongker
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2020/6/14 13:03
 */

package node

import (
	"fmt"
)

type Group struct {
	items map[string]*Node
}

func NewGroup()  *Group{
	return &Group{items: map[string]*Node{}}
}

func (group *Group) Check(project Node) error {
	if group.Has(project.ID) {
		return fmt.Errorf("node:%s is exists", project.ID)
	}

	if group.HasRouter(project.Router) {
		return fmt.Errorf("node router:%s has been used", project.Router)
	}
	return nil
}

// Add add the node to group
func (group *Group) Add(node *Node) {
	group.items[node.ID] = node
}

// Has check node exist
func (group *Group) Has(id string) bool {
	_, ok := group.items[id]
	return ok
}

// HasRouter if node router has been used, return true
func (group *Group) HasRouter(router string) bool {
	for _, item := range group.items {
		if item.Router == router {
			return true
		}
	}

	return false
}

// FindByRouter
func (group *Group) FindByRouter(router string) *Node {
	var project *Node
	for _, item := range group.items {
		if item.Router == router {
			project = item
			break
		}
	}

	return project
}

// Get return node if exist
func (group *Group) Get(id string) *Node {
	return group.items[id]
}

// Delete
func (group *Group) Delete(id string) {
	delete(group.items, id)
}
