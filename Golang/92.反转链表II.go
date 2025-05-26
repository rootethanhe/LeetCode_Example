package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	// 创建虚拟头节点简化边界处理(如left=1时)
	dummy := &ListNode{Next: head}
	pre := dummy

	// 定位到left前驱节点(移动left-1次)
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 头插法翻转区间(right-left次循环)
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next     // 1.保存当前节点的
		cur.Next = next.Next // 2.当前节点指向后续未处理节点
		next.Next = pre.Next // 3.将next插入到pre之后
		pre.Next = next      // 4.更新pre的指向
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Println("->nil")
}

func main() {
	// 示例1：head = [1,2,3,4,5], left = 2, right = 4
	n1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printList(reverseBetween(n1, 2, 4)) // 1->4->3->2->5->nil

	// 示例2：head = [5], left = 1, right = 1
	n2 := &ListNode{5, nil}
	printList(reverseBetween(n2, 1, 1)) // 5->nil
}
