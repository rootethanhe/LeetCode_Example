package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	// 计算链表长度并定位尾节点
	tail := head
	count := 1
	for tail.Next != nil {
		tail = tail.Next
		count++
	}

	// 取模优化: 计算实际有效移动次数
	k %= count
	if k == 0 {
		return head
	}

	// 形成环形链表
	tail.Next = head

	// 定位新尾节点(移动 count - k 次)
	newTail := tail
	for i := 0; i < count-k; i++ {
		newTail = newTail.Next
	}

	// 断开环形链表并返回新头节点
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}

// 辅助函数：将链表转换为字符串输出
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// 构造示例链表1: 1 -> 2 -> 4 -> 5    k值: 2
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next = &ListNode{Val: 5}
	l1K := 2
	printList(rotateRight(l1, l1K))

	// 构造示例链表2: 0 -> 1 -> 2    k值: 4
	l2 := &ListNode{Val: 0}
	l2.Next = &ListNode{Val: 1}
	l2.Next.Next = &ListNode{Val: 2}
	l2K := 4
	printList(rotateRight(l2, l2K))
}
