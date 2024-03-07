package gee

import "strings"

// node 结构表示路由树的节点
type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如[doc, tutorial, intro]
	isWild   bool    // 是否精确匹配， part 含有 ：或 * 时为true
}

// matchChild 在当前节点的子节点中匹配指定的部分
func (n *node) matchChild(part string) *node {
	// 遍历当前节点的子节点
	for _, child := range n.children {
		// 判断子节点是否与指定部分匹配，或者是否为通配符
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 在当前节点的子节点中匹配指定的部分，返回所有匹配的子节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	// 遍历当前节点的子节点
	for _, child := range n.children {
		// 判断子节点是否与指定部分匹配，或者是否为通配符
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 将路由模式插入到路由树中
func (n *node) insert(pattern string, parts []string, height int) {
	// 如果已经到达路由模式的末尾，将当前节点的模式设为要插入的模式
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	// 在当前节点的子节点中匹配当前部分
	child := n.matchChild(part)
	// 如果没有匹配到，创建新的子节点并添加到当前节点的子节点列表中
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	// 递归插入下一层级的路由模式部分
	child.insert(pattern, parts, height+1)
}

// search 在路由树中搜索匹配给定路由部分的节点
func (n *node) search(parts []string, height int) *node {
	// 如果已经到达路由部分的末尾，或当前节点是通配符节点，返回当前节点
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	// 在当前节点的子节点中匹配当前部分
	children := n.matchChildren(part)

	// 递归搜索下一层级的路由部分
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
