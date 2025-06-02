package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

// DFS遍历（前序）
// node: 当前节点
// cur: 当前路径值（根节点到当前节点父节点的数字）
func dfs(node *TreeNode, cur int) int {
	if node == nil {
		return 0
	}

	// 更新路径值：cur×10 + 当前节点值
	cur = cur*10 + node.Val

	// 叶节点：返回当前路径值
	if node.Left == nil && node.Right == nil {
		return cur
	}

	// 非叶节点：递归子树并求和
	return dfs(node.Left, cur) + dfs(node.Right, cur)
}

func main() {
	// 示例1: [1,2,3] → 12+13=25
	root1 := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}
	fmt.Println(sumNumbers(root1)) // 25

	// 示例2: [4,9,0,5,1] → 495+491+40=1026
	root2 := &TreeNode{4,
		&TreeNode{9,
			&TreeNode{5, nil, nil},
			&TreeNode{1, nil, nil},
		},
		&TreeNode{0, nil, nil},
	}
	fmt.Println(sumNumbers(root2)) // 1026

	// 单节点测试
	root3 := &TreeNode{5, nil, nil}
	fmt.Println(sumNumbers(root3)) // 5
}
