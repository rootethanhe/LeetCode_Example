package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.pushLeftBranch(root) // 初始化左链压栈
	return it
}

// 将节点及其左子树压入栈
func (it *BSTIterator) pushLeftBranch(node *TreeNode) {
	for node != nil {
		it.stack = append(it.stack, node)
		node = node.Left
	}
}

func (it *BSTIterator) Next() int {
	// 1. 弹出栈顶节点（当前最小值）
	node := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]

	// 2. 若存在右子树，压入其左链
	if node.Right != nil {
		it.pushLeftBranch(node.Right)
	}

	// 3. 返回节点值
	return node.Val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0 // 栈非空即存在下一个元素
}

func main() {
	// 构建示例树： [7,3,15,null,null,9,20]
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 20}

	// 初始化迭代器
	it := Constructor(root)

	// 验证输出序列
	fmt.Println(it.Next())    // 3
	fmt.Println(it.Next())    // 7
	fmt.Println(it.HasNext()) // true
	fmt.Println(it.Next())    // 9
	fmt.Println(it.HasNext()) // true
	fmt.Println(it.Next())    // 15
	fmt.Println(it.HasNext()) // true
	fmt.Println(it.Next())    // 20
	fmt.Println(it.HasNext()) // false
}
