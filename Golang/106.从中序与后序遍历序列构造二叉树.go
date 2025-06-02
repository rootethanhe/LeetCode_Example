package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	// 创建哈希表加速查找
	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}

	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1, indexMap)
}

func build(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int, indexMap map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}

	// 根节点是后序末位元素
	rootVal := postorder[postEnd]
	root := &TreeNode{Val: rootVal}

	// 获取中序根节点位置
	rootIndex := indexMap[rootVal]
	leftSize := rootIndex - inStart // 左子树节点数

	// 递归构建子树
	root.Left = build(inorder, inStart, rootIndex-1, postorder, postStart, postStart+leftSize-1, indexMap)
	root.Right = build(inorder, rootIndex+1, inEnd, postorder, postStart+leftSize, postEnd-1, indexMap)

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
	// 示例1: inorder=[9,3,15,20,7] postorder=[9,15,7,20,3]
	root1 := buildTree(
		[]int{9, 3, 15, 20, 7},
		[]int{9, 15, 7, 20, 3},
	)
	printTree(root1) // 按层输出: [3,9,20,null,null,15,7]

	// 示例2: inorder=[-1] postorder=[-1]
	root2 := buildTree([]int{-1}, []int{-1})
	printTree(root2) // 输出: [-1]
}
