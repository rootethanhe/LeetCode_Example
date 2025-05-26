package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false // 空树直接返回false[3,5](@ref)
	}
	// 如果是叶子节点，检查当前值是否等于剩余targetSum
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	// 递归检查左右子树，传递更新后的targetSum
	return hasPathSum(root.Left, targetSum-root.Val) ||
		hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	// 示例1：输入 root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
	root1 := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}},
	}
	fmt.Println(hasPathSum(root1, 22)) // 输出: true

	// 示例2：输入 root = [1,2,3], targetSum = 5
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(hasPathSum(root2, 5)) // 输出: false
}
