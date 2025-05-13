package main

import "fmt"

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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// 辅助函数递归判断两棵树是否镜像对称
func isMirror(left, right *TreeNode) bool {
	// 两个节点都为空时对称
	if left == nil && right == nil {
		return true
	}
	// 一个为空一个不为空时不对称
	if left == nil || right == nil {
		return false
	}
	// 节点值必须相等，且子树满足镜像关系
	return left.Val == right.Val &&
		isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}

func main() {
	// 构造二叉树, 对称示例
	treeNode1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	// 进行调试判断
	fmt.Println("Is Symmetric: ", isSymmetric(treeNode1))

	// 构造二叉树, 不对称示例
	treeNode2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	// 进行调试判断
	fmt.Println("Is Symmetric: ", isSymmetric(treeNode2))
}
