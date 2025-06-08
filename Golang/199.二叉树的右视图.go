package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{} // 空树直接返回空切片
	}

	// 初始化队列和结果集
	queue := []*TreeNode{root}
	res := []int{}

	for len(queue) > 0 {
		levelSize := len(queue) // 当前层节点数
		var lastNodeVal int     // 存储当前层最后一个节点值

		// 遍历当前层所有节点
		for i := 0; i < levelSize; i++ {
			node := queue[0]  // 队首出队
			queue = queue[1:] // 更新队列

			// 记录最后一个节点值（每层最后遍历到的节点）
			if i == levelSize-1 {
				lastNodeVal = node.Val
			}

			// 子节点入队（左子树先入队保证层序从左到右）
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 当前层最右节点值加入结果
		res = append(res, lastNodeVal)
	}
	return res
}

func main() {
	// 示例1: [1,2,3,null,5,null,4]
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Right = &TreeNode{Val: 5}
	root1.Right.Right = &TreeNode{Val: 4}
	fmt.Println(rightSideView(root1)) // [1,3,4]

	// 示例2: [1,2,3,4,null,null,null,5]
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Left = &TreeNode{Val: 4}
	root2.Left.Left.Right = &TreeNode{Val: 5}
	fmt.Println(rightSideView(root2)) // [1,3,4,5]

	// 示例3: [1,null,3]
	root3 := &TreeNode{Val: 1}
	root3.Right = &TreeNode{Val: 3}
	fmt.Println(rightSideView(root3)) // [1,3]

	// 示例4: 空树
	var root4 *TreeNode
	fmt.Println(rightSideView(root4)) // []
}
