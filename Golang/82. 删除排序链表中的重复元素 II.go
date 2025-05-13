package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 创建哑节点，处理头节点可能被删除的情况
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	for cur != nil {
		// 检测重复区域：当前节点与下一节点值相同
		if cur.Next != nil && cur.Val == cur.Next.Val {
			// 找到重复区域末尾
			duplicateVal := cur.Val
			for cur != nil && cur.Val == duplicateVal {
				cur = cur.Next
			}
			// 跳过整个重复区域
			pre.Next = cur
		} else {
			// 非重复节点正常移动指针
			pre = cur
			cur = cur.Next
		}
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
	// 示例1：1->2->3->3->4->4->5 → 1->2->5
	n1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
	printList(deleteDuplicates(n1)) // 1->2->5->nil

	// 示例2：1->1->1->2->3 → 2->3
	n2 := &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}
	printList(deleteDuplicates(n2)) // 2->3->nil
}
