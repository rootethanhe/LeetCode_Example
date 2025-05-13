package main

import (
	"fmt"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// 辅助函数：层序打印二叉树
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}

	var result []string
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]  // 取出队列头部节点
		queue = queue[1:] // 移除已处理节点

		result = append(result, fmt.Sprintf("%d", node.Val))

		// 将非空子节点加入队列
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	fmt.Println(strings.Join(result, ","))
}

func main() {
	// 构造示例树
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	// 翻转并验证
	inverted := invertTree(root)
	printTree(inverted) // 应输出层序：4,7,2,9,6,3,1
}
