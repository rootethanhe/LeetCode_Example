package main

import (
	"container/heap"
	"fmt"
)

type minHeap [][3]int // [sum, i, j]

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([3]int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	// 初始化堆
	h := &minHeap{}
	heap.Init(h)
	// 初始候选：nums1 前 min(k, len(nums1)) 个元素与 nums2[0] 组合
	n1 := min(k, len(nums1))
	for i := 0; i < n1; i++ {
		heap.Push(h, [3]int{nums1[i] + nums2[0], i, 0})
	}

	// 收集结果
	result := [][]int{}
	for h.Len() > 0 && len(result) < k {
		top := heap.Pop(h).([3]int)
		i, j := top[1], top[2]
		result = append(result, []int{nums1[i], nums2[j]})

		// 扩展：同一 nums1[i] 的下一个 nums2 元素
		if j+1 < len(nums2) {
			heap.Push(h, [3]int{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 示例 1
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k)) // 输出: [[1,2] [1,4] [1,6]]

	// 示例 2
	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2
	fmt.Println(kSmallestPairs(nums1, nums2, k)) // 输出: [[1,1] [1,1]]
}
