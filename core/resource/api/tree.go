package api

import (
	"fmt"
	"github.com/ebar-go/ego/utils/number"
	"github.com/pkg/errors"
)

// upstream 路由节点
type node struct {
	// 保存这个节点上的URL路径
	// 例如上图中的search和support, 共同的parent节点的path="s"
	// 后面两个节点的path分别是"earch"和"upport"
	path string

	// 判断当前节点路径是不是参数节点, 例如上图的:post部分就是wildChild节点
	wildChild bool

	// 节点类型包括static, root, param, catchAll
	// static: 静态节点, 例如上面分裂出来作为parent的s
	// root: 如果插入的节点是第一个, 那么是root节点
	// catchAll: 有*匹配的节点
	// param: 除上面外的节点
	nType nodeType

	// 记录路径上最大参数个数
	maxParams uint8

	// 和children[]对应, 保存的是分裂的分支的第一个字符
	// 例如search和support, 那么s节点的indices对应的"eu"
	// 代表有两个分支, 分支的首字母分别是e和u
	indices string

	// 保存孩子节点
	children []*node

	// 当前节点的匹配目标
	handle Handle

	// 优先级, 看起来没什么卵用的样子@_@
	priority uint32
}

type nodeType uint8

const (
	static nodeType = iota // default
	root
	param
	catchAll
)

func countParams(path string) uint8 {
	var n uint
	for i := 0; i < len(path); i++ {
		if path[i] != ':' && path[i] != '*' {
			continue
		}
		n++
	}
	if n >= 255 {
		return 255
	}
	return uint8(n)
}

func (n *node) addRoute(path string, handle Handle) error {
	fullPath := path
	n.priority++
	numParams := countParams(path)

	if len(n.path) > 0 || len(n.children) > 0 {
	walk:
		for {
			// Update maxParams of the current upstream
			// 更新当前node的最大参数个数
			if numParams > n.maxParams {
				n.maxParams = numParams
			}

			// Find the longest common prefix.
			// This also implies that the common prefix contains no ':' or '*'
			// since the existing key can't contain those chars.
			// 找到最长公共前缀
			i := 0
			max := number.Min(len(path), len(n.path))
			// 匹配相同的字符
			for i < max && path[i] == n.path[i] {
				i++
			}

			// Split edge
			// 说明前面有一段是匹配的, 例如之前为:/search,现在来了一个/support
			// 那么会将/s拿出来作为parent节点, 将child节点变成earch和upport
			if i < len(n.path) {
				// 将原本路径的i后半部分作为前半部分的child节点
				child := node{
					path:      n.path[i:],
					wildChild: n.wildChild,
					nType:     static,
					indices:   n.indices,
					children:  n.children,
					handle:    n.handle,
					priority:  n.priority - 1,
				}

				// Update maxParams (max of all children)
				// 更新最大参数个数
				for i := range child.children {
					if child.children[i].maxParams > child.maxParams {
						child.maxParams = child.children[i].maxParams
					}
				}
				// 当前节点的孩子节点变成刚刚分出来的这个后半部分节点
				n.children = []*node{&child}
				// []byte for proper unicode char conversion, see #65
				n.indices = string([]byte{n.path[i]})
				// 路径变成前i半部分path
				n.path = path[:i]
				n.handle = nil
				n.wildChild = false
			}

			// Make new upstream a child of this upstream
			// 同时, 将新来的这个节点插入新的parent节点中当做孩子节点
			if i < len(path) {
				// i的后半部分作为路径, 即上面例子support中的upport
				path = path[i:]

				// 如果n是参数节点(包含:或者*)
				if n.wildChild {
					n = n.children[0]
					n.priority++

					// Update maxParams of the child upstream
					if numParams > n.maxParams {
						n.maxParams = numParams
					}
					numParams--

					// Check if the wildcard matches
					// 例如: /blog/:ppp 和 /blog/:ppppppp, 需要检查更长的通配符
					if len(path) >= len(n.path) && n.path == path[:len(n.path)] {
						// check for longer wildcard, e.g. :name and :names
						if len(n.path) >= len(path) || path[len(n.path)] == '/' {
							continue walk
						}
					}

					panic("path segment '" + path +
						"' conflicts with existing wildcard '" + n.path +
						"' in path '" + fullPath + "'")
				}

				c := path[0]

				// slash after param
				if n.nType == param && c == '/' && len(n.children) == 1 {
					n = n.children[0]
					n.priority++
					continue walk
				}

				// Check if a child with the next path byte exists
				// 检查路径是否已经存在, 例如search和support第一个字符相同
				for i := 0; i < len(n.indices); i++ {
					// 找到第一个匹配的字符
					if c == n.indices[i] {
						i = n.incrementChildPrio(i)
						n = n.children[i]
						continue walk
					}
				}

				// Otherwise insert it
				// new一个node
				if c != ':' && c != '*' {
					// []byte for proper unicode char conversion, see #65
					// 记录第一个字符,并放在indices中
					n.indices += string([]byte{c})
					child := &node{
						maxParams: numParams,
					}
					// 增加孩子节点
					n.children = append(n.children, child)
					n.incrementChildPrio(len(n.indices) - 1)
					n = child
				}
				// 插入节点
				n.insertChild(numParams, path, fullPath, handle)
				return nil

				// 说明是相同的路径,仅仅需要将handle替换就OK
				// 如果是nil那么说明取消这个handle, 不是空不允许
			} else if i == len(path) { // Make upstream a (in-path) leaf
				if n.handle != nil && handle != nil {
					return errors.New("a handle is already registered for path '" + fullPath + "'")
				}
				n.handle = handle
			}
			return nil
		}
	} else { // Empty tree
		// 如果是空树, 那么插入节点
		n.insertChild(numParams, path, fullPath, handle)
		// 节点的种类是root
		n.nType = root
	}

	return nil
}

// increments priority of the given child and reorders if necessary
func (n *node) incrementChildPrio(pos int) int {
	n.children[pos].priority++
	prio := n.children[pos].priority

	// adjust position (move to front)
	newPos := pos
	for newPos > 0 && n.children[newPos-1].priority < prio {
		// swap upstream positions
		n.children[newPos-1], n.children[newPos] = n.children[newPos], n.children[newPos-1]

		newPos--
	}

	// build new index char string
	if newPos != pos {
		n.indices = n.indices[:newPos] + // unchanged prefix, might be empty
			n.indices[pos:pos+1] + // the index char we move
			n.indices[newPos:pos] + n.indices[pos+1:] // rest without char at 'pos'
	}

	return newPos
}

func (n *node) insertChild(numParams uint8, path, fullPath string, handle Handle) {
	var offset int // already handled bytes of the path

	// 找到前缀, 直到遇到第一个wildcard匹配的参数
	for i, max := 0, len(path); numParams > 0; i++ {
		c := path[i]
		if c != ':' && c != '*' {
			continue
		}

		// find wildcard end (either '/' or path end)
		end := i + 1
		// 下面判断:或者*之后不能再有*或者:, 这样是属于参数错误
		// 除非到了下一个/XXX
		for end < max && path[end] != '/' {
			switch path[end] {
			// the wildcard name must not contain ':' and '*'
			case ':', '*':
				panic("only one wildcard per path segment is allowed, has: '" +
					path[i:] + "' in path '" + fullPath + "'")
			default:
				end++
			}
		}

		// check if this Node existing children which would be
		// unreachable if we insert the wildcard here
		if len(n.children) > 0 {
			panic("wildcard route '" + path[i:end] +
				"' conflicts with existing children in path '" + fullPath + "'")
		}

		// check if the wildcard has a name
		// 下面的判断说明只有:或者*,没有name,这也是不合法的
		if end-i < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}

		// 如果是':',那么匹配一个参数
		if c == ':' { // param
			// split path at the beginning of the wildcard
			// 节点path是参数前面那么一段, offset代表已经处理了多少path中的字符
			if i > 0 {
				n.path = path[offset:i]
				offset = i
			}
			// 构造一个child
			child := &node{
				nType:     param,
				maxParams: numParams,
			}
			n.children = []*node{child}
			n.wildChild = true
			// 下次的循环就是这个新的child节点了
			n = child
			// 最长匹配, 所以下面节点的优先级++
			n.priority++
			numParams--

			// if the path doesn't end with the wildcard, then there
			// will be another non-wildcard subpath starting with '/'
			if end < max {
				n.path = path[offset:end]
				offset = end

				child := &node{
					maxParams: numParams,
					priority:  1,
				}
				n.children = []*node{child}
				n = child
			}

		} else { // catchAll
			// *匹配所有参数
			if end != max || numParams > 1 {
				panic("catch-all routes are only allowed at the end of the path in path '" + fullPath + "'")
			}

			if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
				panic("catch-all conflicts with existing handle for the path segment root in path '" + fullPath + "'")
			}

			// currently fixed width 1 for '/'
			i--
			fmt.Println(i, path)
			if path[i] != '/' {
				panic("no / before catch-all in path '" + fullPath + "'")
			}

			n.path = path[offset:i]

			// first upstream: catchAll upstream with empty path
			child := &node{
				wildChild: true,
				nType:     catchAll,
				maxParams: 1,
			}
			n.children = []*node{child}
			n.indices = string(path[i])
			n = child
			n.priority++

			// second upstream: upstream holding the variable
			child = &node{
				path:      path[i:],
				nType:     catchAll,
				maxParams: 1,
				handle:    handle,
				priority:  1,
			}
			n.children = []*node{child}

			return
		}
	}

	// insert remaining path part and handle to the leaf
	n.path = path[offset:]
	n.handle = handle
}

// Returns the handle registered with the given path (key). The values of
// wildcards are saved to a map.
// If no handle can be found, a TSR (trailing slash redirect) recommendation is
// made if a handle exists with an extra (without the) trailing slash for the
// given path.
func (n *node) getValue(path string) (handle Handle) {
walk: // outer loop for walking the tree
	for {
		// 意思是如果还没有走到路径end
		if len(path) > len(n.path) {
			// 前面一段必须和当前节点的path一样才OK
			if path[:len(n.path)] == n.path {
				path = path[len(n.path):]
				// If this upstream does not have a wildcard (param or catchAll)
				// child,  we can just look up the next child upstream and continue
				// to walk down the tree
				// 如果不是参数节点, 那么根据分支walk到下一个节点就OK
				if !n.wildChild {
					c := path[0]
					// 找到分支的第一个字符=>找到child
					for i := 0; i < len(n.indices); i++ {
						if c == n.indices[i] {
							n = n.children[i]
							continue walk
						}
					}

					// Nothing found.
					// We can recommend to redirect to the same URL without a
					// trailing slash if a leaf exists for that path.
					// tsr = (path == "/" && n.handle != nil)
					return

				}

				// handle wildcard child
				// 下面处理通配符参数节点
				n = n.children[0]
				switch n.nType {
				// 如果是普通':'节点, 那么找到/或者path end, 获得参数
				case param:
					// find param end (either '/' or path end)
					end := 0
					for end < len(path) && path[end] != '/' {
						end++
					}
					// 获取参数
					// save param value
					/*
						if p == nil {
							// lazy allocation
							p = make(Params, 0, n.maxParams)
						}
						i := len(p)
						p = p[:i+1] // expand slice within preallocated capacity
						// 获取key和value
						p[i].Key = n.path[1:]
						p[i].Value = path[:end]
					*/

					// we need to go deeper!
					// 如果参数还没处理完, 继续walk
					if end < len(path) {
						if len(n.children) > 0 {
							path = path[end:]
							n = n.children[0]
							continue walk
						}

						// ... but we can't
						// tsr = (len(path) == end+1)
						return
					}
					// 否则获得handle返回就OK
					if handle = n.handle; handle != nil {
						return
					} else if len(n.children) == 1 {
						// No handle found. Check if a handle for this path + a
						// trailing slash exists for TSR recommendation
						n = n.children[0]
						// tsr = (n.path == "/" && n.handle != nil)
					}

					return

				case catchAll:

					handle = n.handle
					return

				default:
					panic("invalid upstream type")
				}
			}
			// 走到路径end
		} else if path == n.path {
			// We should have reached the upstream containing the handle.
			// Check if this upstream has a handle registered.
			// 判断这个路径节点是都存在handle, 如果存在, 那么就可以直接返回了.
			if handle = n.handle; handle != nil {
				return
			}
			// 下面判断是不是需要进入重定向
			if path == "/" && n.wildChild && n.nType != root {
				// tsr = true
				return
			}

			// No handle found. Check if a handle for this path + a
			// trailing slash exists for trailing slash recommendation
			// 判断path+'/'是否存在handle
			for i := 0; i < len(n.indices); i++ {
				if n.indices[i] == '/' {
					n = n.children[i]
					// tsr = (len(n.path) == 1 && n.handle != nil) ||
					// (n.nType == catchAll && n.children[0].handle != nil)
					return
				}
			}

			return
		}

		// Nothing found. We can recommend to redirect to the same URL with an
		// extra trailing slash if a leaf exists for that path
		/*
			tsr = (path == "/") ||
				(len(n.path) == len(path)+1 && n.path[len(path)] == '/' &&
					path == n.path[:len(n.path)-1] && n.handle != nil)
		*/
		return
	}
}
