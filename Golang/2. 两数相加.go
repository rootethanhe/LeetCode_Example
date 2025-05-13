package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 哑节点简化操作
	curr := dummy
	carry := 0

	// 循环直到所有节点处理完且无进位
	for l1 != nil || l2 != nil || carry > 0 {
		sumVal := carry // 初始化为进位值

		// 处理当前位相加
		if l1 != nil {
			sumVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sumVal += l2.Val
			l2 = l2.Next
		}

		// 计算新节点值和进位
		carry = sumVal / 10
		curr.Next = &ListNode{Val: sumVal % 10}
		curr = curr.Next
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
	// l1 = [2,4,3], l2 = [5,6,4]
	printList(addTwoNumbers(&ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}, &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}))

	// l1 = [0], l2 = [0]
	printList(addTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 0}))

	// l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	printList(addTwoNumbers(&ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}, &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}))
}
