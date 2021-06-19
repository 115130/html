package gee

import (
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	childern []*node //子节点
	isWild   bool    //是否精确匹配
}

//第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.childern {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.childern {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//节点的插入，递归查找每一层节点，没有匹配到则插入
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.childern = append(n.childern, child)
	}
	child.insert(pattern, parts, height+1)
}

//节点的查询
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {//判断是否匹配成功
			return nil
		}
		return n
	}

	part := parts[height]
	childern := n.matchChildren(part)

	for _, child := range childern {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
