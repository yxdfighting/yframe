package yframe

import "strings"

//定义和实现索引前缀树
//主要用于路由的动态前缀匹配
//可以与gin由相同的匹配规则
//yxdtest/:user/helloworld     对yxdtest/8025/helloworld    错yxdtest/8026/111/helloworld
//yxdtest/*user/helloworld	   对yxdtest/8025    对yxdtest/helloworld/8026/111/helloworld
type node struct{
	pattern string //从头到当前节点的pattern
	part    string //当前节点
	children []*node
	isFuzzy   bool //表示当前节点是否是精确匹配，比如： * 此时该值为true
}

//查找当前node可以满足input的一个子node
func (n *node) matchChild(input string) *node{
	for _,v := range n.children{
		if v.part == input || v.isFuzzy{
			return v
		}
	}
	return nil
}

//查找当前node可以满足input的所有子node列表
func (n *node) matchChildren (input string) []*node{
	var res []*node
	for _,v := range n.children{
		if v.part == input || v.isFuzzy{
			res = append(res,v)
		}
	}
	return res
}

//插入逻辑：
//递归实现，当trie层树与parts元素数量相同时，递归终止
//对于每一个节点，part应该是对应当前层高的parts[idx]元素；children表示所有满足要求的child列表集合
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isFuzzy: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}





