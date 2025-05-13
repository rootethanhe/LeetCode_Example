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

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	average := make([]float64, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0.0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += float64(node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		average = append(average, sum/float64(levelSize))
	}

	return average
}

func main() {
	// 构造示例  root = [3,9,20,null,null,15,7]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// 求层平均值 [3.00000,14.50000,11.00000]
	fmt.Println("Average Of Levels: ", averageOfLevels(root))
}
