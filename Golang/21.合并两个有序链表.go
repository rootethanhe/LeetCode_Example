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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // 创建哑节点简化边界处理
	current := dummy

	// 双指针遍历两个链表
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// 处理剩余链表部分
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
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
	// 构造示例链表1: 1 -> 2 -> 4
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}

	// 构造示例链表2: 1 -> 3 -> 4
	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	// 合并链表
	merged := mergeTwoLists(l1, l2)

	// 打印合并结果
	fmt.Print("合并结果: ")
	printList(merged) // 应输出: 1 -> 1 -> 2 -> 3 -> 4 -> 4 -> nil
}
