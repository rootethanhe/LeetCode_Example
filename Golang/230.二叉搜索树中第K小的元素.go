package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{} // 初始化栈

	for len(stack) > 0 || root != nil {
		// 1. 左子树入栈（深度优先）
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 2. 弹出当前最小节点
		n := len(stack)
		node := stack[n-1]
		stack = stack[:n-1]

		// 3. 检查是否第K小
		k--
		if k == 0 {
			return node.Val // 找到目标
		}

		// 4. 转向右子树
		root = node.Right
	}
	return -1 // k超出范围
}

func main() {
	// 示例1：root = [3,1,4,null,2], k=1 → 输出1
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 2}
	fmt.Println(kthSmallest(root1, 1)) // 1

	// 示例2：root = [5,3,6,2,4,null,null,1], k=3 → 输出3
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 6}
	root2.Left.Left = &TreeNode{Val: 2}
	root2.Left.Right = &TreeNode{Val: 4}
	root2.Left.Left.Left = &TreeNode{Val: 1}
	fmt.Println(kthSmallest(root2, 3)) // 3
}
