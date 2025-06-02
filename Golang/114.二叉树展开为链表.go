package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			// 1. 找左子树的最右节点
			predecessor := curr.Left
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}

			// 2. 连接原右子树
			predecessor.Right = curr.Right

			// 3. 左子树移到右侧
			curr.Right = curr.Left
			curr.Left = nil
		}
		// 4. 处理下一个节点
		curr = curr.Right
	}
}

// 辅助函数：打印链表
func printList(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	curr := root
	for curr != nil {
		fmt.Printf("%d", curr.Val)
		if curr.Right != nil {
			fmt.Print("->")
		}
		curr = curr.Right
	}
	fmt.Println()
}

func main() {
	// 示例1: [1,2,5,3,4,null,6]
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			&TreeNode{4, nil, nil},
		},
		&TreeNode{5,
			nil,
			&TreeNode{6, nil, nil},
		},
	}
	flatten(root)
	printList(root) // 输出: 1->2->3->4->5->6

	// 示例2: 空树
	var root2 *TreeNode
	flatten(root2)
	printList(root2) // 输出: []

	// 示例3: 单节点
	root3 := &TreeNode{0, nil, nil}
	flatten(root3)
	printList(root3) // 输出: 0
}
