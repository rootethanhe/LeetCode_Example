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

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, left, right int) *TreeNode {
	if left > right { // 递归终止条件：子数组为空
		return nil
	}
	mid := left + (right-left)/2 // 取中间元素作为根节点(向下取整)
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildBST(nums, left, mid-1)   // 递归构建左子树
	root.Right = buildBST(nums, mid+1, right) // 递归构建右子树
	return root
}

// 打印二叉树结构（层序遍历）
func printTree(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}

	result := []any{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				result = append(result, nil)
				continue
			}
			result = append(result, node.Val)

			// 将子节点加入队列（包括空子节点）
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// 去除末尾连续的 null
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func main() {
	// 示例1：输入 nums = [-10,-3,0,5,9]
	nums1 := []int{-10, -3, 0, 5, 9}
	root1 := sortedArrayToBST(nums1)
	fmt.Println(printTree(root1))
	// 可能的输出结构：[0,-3,9,-10,null,5]

	// 示例2：输入 nums = [1,3]
	nums2 := []int{1, 3}
	root2 := sortedArrayToBST(nums2)
	fmt.Println(printTree(root2))
	// 可能的输出结构：[3,1] 或 [1,null,3]
}
