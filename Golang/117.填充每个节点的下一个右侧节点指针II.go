package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	current := root // 当前层起始节点

	for current != nil {
		dummy := &Node{} // 哑节点（下一层链表头的前驱）
		prev := dummy    // 下一层连接指针

		// 遍历当前层所有节点
		for current != nil {
			// 连接左子节点
			if current.Left != nil {
				prev.Next = current.Left
				prev = prev.Next // 移动连接指针
			}
			// 连接右子节点
			if current.Right != nil {
				prev.Next = current.Right
				prev = prev.Next // 移动连接指针
			}
			current = current.Next // 横向移动到同层下一节点
		}

		current = dummy.Next // 进入下一层（哑节点指向下一层头）
	}
	return root
}

// 辅助函数：按层打印Next连接
func printTree(root *Node) {
	current := root
	for current != nil {
		node := current
		for node != nil {
			fmt.Printf("%d", node.Val)
			if node.Next != nil {
				fmt.Print("->")
			}
			node = node.Next
		}
		fmt.Println("->nil")
		current = current.Left // 下一层首节点在左子
	}
}

func main() {

	// 示例1: [1,2,3,4,5,null,7]
	root := &Node{
		1,
		&Node{
			2,
			&Node{4, nil, nil, nil},
			&Node{5, nil, nil, nil},
			nil,
		},
		&Node{
			3,
			nil,
			&Node{7, nil, nil, nil},
			nil},
		nil,
	}

	connect(root)
	printTree(root) // 输出: 1->nil, 2->3->nil, 4->5->7->nil

	// 示例2: 空树
	root2 := connect(nil)
	fmt.Println(root2) // 输出: nil

	// 示例3: [1,2,3,4,null,null,5]
	root3 := &Node{
		1,
		&Node{
			2,
			&Node{4, nil, nil, nil},
			nil,
			nil,
		},
		&Node{
			3,
			nil,
			&Node{5, nil, nil, nil},
			nil},
		nil,
	}
	connect(root3)
	printTree(root3) // 输出: 1->nil, 2->3->nil, 4->5->nil
}
