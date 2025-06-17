package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	return dfs(node, visited)
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	// 已克隆则直接返回
	if clone, exists := visited[node]; exists {
		return clone
	}

	// 创建新节点并登记
	cloneNode := &Node{Val: node.Val}
	visited[node] = cloneNode

	// 递归克隆邻居
	for _, neighbor := range node.Neighbors {
		cloneNeighbor := dfs(neighbor, visited)
		cloneNode.Neighbors = append(cloneNode.Neighbors, cloneNeighbor)
	}

	return cloneNode
}

func main() {
	// 示例1：四节点图
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}
	cloned := cloneGraph(node1)
	fmt.Println(cloned.Val) // 输出: 1

	// 示例2：单节点空邻居
	singleNode := &Node{Val: 1}
	clonedSingle := cloneGraph(singleNode)
	fmt.Println(clonedSingle.Neighbors) // 输出: []

	// 示例3：空图
	fmt.Println(cloneGraph(nil)) // 输出: nil
}
