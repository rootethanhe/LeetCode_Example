package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// 递归终止条件: 空链表或单节点
	if head == nil || head.Next == nil {
		return head
	}

	// Step 1: 找到链表中点并分割
	mid := getMid(head)
	left := head
	right := mid.Next
	mid.Next = nil // 断开左右链表

	// Step 2: 递归排序左右子链表
	left = sortList(left)
	right = sortList(right)

	// Step 3: 合并有序链表
	return merge(left, right)
}

// 快慢指针找中点(慢指针停在中间左侧)
func getMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 合并两个有序链表
func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 虚拟头结点简化操作
	tail := dummy

	// 比较节点值，按升序连接
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	// 处理剩余节点
	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummy.Next
}

// 辅助函数：打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d → ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// 示例1：输入 [4,2,1,3] → 输出 [1,2,3,4]
	head1 := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	sorted1 := sortList(head1)
	printList(sorted1) // 输出：1 → 2 → 3 → 4

	// 示例2：输入 [-1,5,3,4,0] → 输出 [-1,0,3,4,5]
	head2 := &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}}
	sorted2 := sortList(head2)
	printList(sorted2) // 输出：-1 → 0 → 3 → 4 → 5

	// 示例3：输入 [] → 输出 []
	var head3 *ListNode
	sorted3 := sortList(head3)
	printList(sorted3) // 输出：空
}
