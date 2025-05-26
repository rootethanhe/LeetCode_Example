package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// 创建两个虚拟头节点，用于分别构建小于x和大于等于x的链表
	lessHead := &ListNode{}
	greaterHead := &ListNode{}
	lessTail := lessHead       // 跟踪小于x链表的尾部
	greaterTail := greaterHead // 跟踪大于等于x链表的尾部

	// 遍历原链表，分配节点到对应子链表
	for head != nil {
		if head.Val < x {
			lessTail.Next = head
			lessTail = lessTail.Next
		} else {
			greaterTail.Next = head
			greaterTail = greaterTail.Next
		}
		head = head.Next // 移动到下一个节点
	}

	// 处理大于等于链表的尾部，断开可能的循环引用
	greaterTail.Next = nil

	// 合并两个链表
	lessTail.Next = greaterHead.Next

	// 返回合并后的头节点
	return lessHead.Next
}

// 辅助函数: 打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// 构建输入链表：1 -> 4 -> 3 -> 2 -> 5 -> 2
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 2}
	l1.Next.Next.Next.Next = &ListNode{Val: 5}
	l1.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	printList(partition(l1, 3)) // 输出：1->2->2->4->3->5->nil

	// 输入链表：2 -> 1
	l2 := &ListNode{Val: 2}
	l2.Next = &ListNode{Val: 1}
	printList(partition(l2, 2)) // 输出：1->2->nil
}
