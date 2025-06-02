package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)   // 始终向左遍历计算高度
	rightHeight := getHeight(root.Right) // 始终向左遍历计算高度

	if leftHeight == rightHeight {
		// 左子树是满二叉树 + 根节点 + 递归右子树
		return (1 << leftHeight) + countNodes(root.Right)
	}
	// 右子树是满二叉树 + 根节点 + 递归左子树
	return (1 << rightHeight) + countNodes(root.Left)
}

// 辅助函数：向左遍历计算高度
func getHeight(node *TreeNode) int {
	height := 0
	for node != nil {
		height++
		node = node.Left
	}
	return height
}

func main() {
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 5}
	root1.Right.Left = &TreeNode{Val: 6}
	fmt.Println(countNodes(root1)) // 6

	fmt.Println(countNodes(nil)) // 0

	root3 := &TreeNode{Val: 1}
	fmt.Println(countNodes(root3)) // 1
}
