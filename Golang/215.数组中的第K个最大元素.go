package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	target := n - k // 目标索引（升序序列中的位置）

	var quickSelect func(left, right int) int
	quickSelect = func(left, right int) int {
		if left == right {
			return nums[left]
		}

		// 随机选择基准值并置于起始位置
		pivotIndex := left + rand.Intn(right-left+1)
		nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
		pivot := nums[left]

		// 初始化三区指针
		lt, gt, i := left, right, left+1

		// 三向划分
		for i <= gt {
			switch {
			case nums[i] < pivot:
				nums[lt], nums[i] = nums[i], nums[lt]
				lt++
				i++
			case nums[i] > pivot:
				nums[i], nums[gt] = nums[gt], nums[i]
				gt--
			default:
				i++
			}
		}

		// 根据目标位置决策
		switch {
		case target < lt:
			return quickSelect(left, lt-1)
		case target > gt:
			return quickSelect(gt+1, right)
		default:
			return pivot
		}
	}

	return quickSelect(0, n-1)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for i := 1; i <= 5; i++ {
		k := i
		result := findKthLargest(nums, k)
		fmt.Printf("第%d大元素 = %d\n", k, result)
	}
	fmt.Println()

	// 大量重复元素测试
	nums = make([]int, 1000)
	for i := range nums {
		nums[i] = 1 // 全1数组
	}
	nums[0], nums[1], nums[2] = 3, 2, 4 // 添加少量不同值

	fmt.Println("全1数组中：")
	fmt.Println("第1大元素 =", findKthLargest(nums, 1)) // 4
	fmt.Println("第2大元素 =", findKthLargest(nums, 2)) // 3
	fmt.Println("第3大元素 =", findKthLargest(nums, 3)) // 2
	fmt.Println("第4大元素 =", findKthLargest(nums, 4)) // 1
}
