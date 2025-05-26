package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head} // 哑节点简化头结点删除逻辑
	fast, slow := dummy, dummy

	// 快指针先移动n步
	for i := 0; i <= n; i++ { // 多走一步，使慢指针停在目标节点的前驱位置
		fast = fast.Next
	}

	// 同步移动快指针直到链表末尾
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 删除目标节点
	slow.Next = slow.Next.Next // 直接跳过目标节点
	return dummy.Next          // 返回哑节点的下一个节点作为新头结点
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
	// 构造示例链表1: 1 -> 2 -> 3 -> 4 -> 5    n值: 2  输出: [1,2,3,5]
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next.Next = &ListNode{Val: 5}
	l1N := 2
	printList(removeNthFromEnd(l1, l1N))

	// 构造示例链表2: 1    n值: 1    输出: []
	l2 := &ListNode{Val: 1}
	l2N := 1
	printList(removeNthFromEnd(l2, l2N))

	// 构造示例链表3: 1 -> 2    n值: 1    输出: [1]
	l3 := &ListNode{Val: 1}
	l3.Next = &ListNode{Val: 2}
	l3N := 1
	printList(removeNthFromEnd(l3, l3N))
}
