package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 同时为空：结构相同
	if p == nil && q == nil {
		return true
	}
	// 仅一个为空：结构不同
	if p == nil || q == nil {
		return false
	}
	// 节点值不同：数值不同
	if p.Val != q.Val {
		return false
	}
	// 递归比较左右子树
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	// 示例1：p = [1,2,3], q = [1,2,3]
	p1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(isSameTree(p1, q1)) // true

	// 示例2：p = [1,2], q = [1,null,2]
	p2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	fmt.Println(isSameTree(p2, q2)) // false

	// 示例3：p = [1,2,1], q = [1,1,2]
	p3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
	q3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	fmt.Println(isSameTree(p3, q3)) // false
}
