package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{} // 空树直接返回
	}

	// 初始化队列和结果集
	queue := []*TreeNode{root}
	res := [][]int{}

	for len(queue) > 0 {
		levelSize := len(queue) // 当前层节点数
		levelVals := make([]int, 0, levelSize)

		// 遍历当前层节点
		for i := 0; i < levelSize; i++ {
			// 队首节点出队
			node := queue[0]
			queue = queue[1:]

			// 收集节点值
			levelVals = append(levelVals, node.Val)

			// 子节点入队（下一层）
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 当前层结果加入总结果
		res = append(res, levelVals)
	}
	return res
}

func main() {
	// 示例1: [3,9,20,null,null,15,7]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}
	fmt.Println(levelOrder(root1)) // [[3] [9 20] [15 7]]

	// 示例2: 单节点
	root2 := &TreeNode{Val: 1}
	fmt.Println(levelOrder(root2)) // [[1]]

	// 示例3: 空树
	var root3 *TreeNode
	fmt.Println(levelOrder(root3)) // []
}
