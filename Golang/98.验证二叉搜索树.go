package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var prev int
	hasPrev := false // 标记是否已访问节点

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true // 空树视为有效BST
		}

		// 1. 遍历左子树
		if !dfs(node.Left) {
			return false
		}

		// 2. 检查当前节点
		if hasPrev && node.Val <= prev {
			return false // 违反递增规则
		}
		prev = node.Val // 更新前驱值
		hasPrev = true  // 标记已访问节点

		// 遍历右子树
		return dfs(node.Right)
	}

	return dfs(root)
}

func main() {
	// 示例1：有效BST [2,1,3]
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}
	fmt.Println(isValidBST(root1)) // true

	// 示例2：无效BST [5,1,4,null,null,3,6]
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 4}
	root2.Right.Left = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 6}
	fmt.Println(isValidBST(root2)) // false

	// 边界案例：单节点树
	root3 := &TreeNode{Val: 1}
	fmt.Println(isValidBST(root3)) // true

	// 边界案例：含int最小值
	root4 := &TreeNode{Val: math.MinInt32}
	root4.Right = &TreeNode{Val: math.MaxInt32}
	fmt.Println(isValidBST(root4)) // true
}
