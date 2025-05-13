package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢指针走一步
		fast = fast.Next.Next // 快指针走两步

		if slow == fast { // 相遇说明有环
			return true
		}
	}

	return false
}

func main() {
	// 构造示例链表 [3 -> 2 -> 0 -> -4]，尾节点连接到索引1的位置（值为2的节点）
	node4 := &ListNode{Val: -4}
	node3 := &ListNode{Val: 0, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 3, Next: node2}
	node4.Next = node2 // 形成环：-4 -> 2

	// 测试环检测
	fmt.Println("Has cycle:", hasCycle(node1)) // 应输出 true

	// 构造无环链表测试
	node4.Next = nil // 断开环
	fmt.Println("Has cycle:", hasCycle(node1))
}
