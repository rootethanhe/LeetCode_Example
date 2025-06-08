package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{} // 空树直接返回空结果
	}

	// 初始化队列和结果集
	queue := []*TreeNode{root}
	var res [][]int
	level := 0 // 层数标记(0开始)

	for len(queue) > 0 {
		levelSize := len(queue) // 当前层节点数
		levelVals := make([]int, 0, levelSize)

		// 遍历当前层节点
		for i := 0; i < levelSize; i++ {
			node := queue[0]  // 队首出队
			queue = queue[1:] // 更新队列

			levelVals = append(levelVals, node.Val) // 存储节点值

			// 子节点入队(下一层)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 奇数层反转节点值列表
		if level%2 == 1 {
			for i, j := 0, len(levelVals)-1; i < j; i, j = i+1, j-1 {
				levelVals[i], levelVals[j] = levelVals[j], levelVals[i]
			}
		}

		res = append(res, levelVals) // 加入结果集
		level++                      // 层数递增
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
	fmt.Println(zigzagLevelOrder(root1)) // [[3] [20 9] [15 7]]

	// 示例2: 单节点
	root2 := &TreeNode{Val: 1}
	fmt.Println(zigzagLevelOrder(root2)) // [[1]]

	// 示例3: 空树
	var root3 *TreeNode
	fmt.Println(zigzagLevelOrder(root3)) // []
}
