package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 构建中序索引映射
	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}

	return helper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, indexMap)
}

func helper(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, indexMap map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	// 前序首元素为根节点
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	// 获取中序根节点位置
	rootIndex := indexMap[rootVal]
	leftNum := rootIndex - inStart // 左子树节点数

	// 递归构建子树
	root.Left = helper(preorder, preStart+1, preStart+leftNum, inorder, inStart, rootIndex-1, indexMap)
	root.Right = helper(preorder, preStart+leftNum+1, preEnd, inorder, rootIndex+1, inEnd, indexMap)

	return root
}

// 按层打印二叉树, 测试使用
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	queue := []*TreeNode{root}
	var res []interface{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, nil)
		} else {
			res = append(res, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}

	// 移除末尾 nil
	i := len(res) - 1
	for i >= 0 && res[i] == nil {
		i--
	}
	fmt.Println(res[:i+1])
}

func main() {
	// 示例1: pre=[3,9,20,15,7], in=[9,3,15,20,7]
	root1 := buildTree(
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
	)
	printTree(root1) // [3 9 20 nil nil 15 7]

	// 示例2: pre=[-1], in=[-1]
	root2 := buildTree([]int{-1}, []int{-1})
	printTree(root2) // [-1]
}
