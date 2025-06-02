package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 终止条件: 遇到nil/p/q直接返回
	if root == nil || root == p || root == q {
		return root
	}

	// 递归搜索左右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 结果合并判断
	switch {
	case left != nil && right != nil: // 左右均找到 --> 当前为LCA
		return root
	case left != nil:
		return left // 仅左子树找到
	default:
		return right // 仅右子树找到或均未找到
	}
}

func main() {
	// 示例1：LCA(5,1)=3
	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}
	node1 := &TreeNode{Val: 1}
	node3.Left, node3.Right = node5, node1
	fmt.Println(lowestCommonAncestor(node3, node5, node1).Val) // 3

	// 示例2：LCA(5,4)=5
	node4 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2, Right: node4}
	node5.Left = &TreeNode{Val: 6}
	node5.Right = node2
	fmt.Println(lowestCommonAncestor(node3, node5, node4).Val) // 5

	// 示例3：LCA(1,2)=1
	node1.Left = &TreeNode{Val: 2}
	fmt.Println(lowestCommonAncestor(node1, node1, node1.Left).Val) // 1
}
