package main

import "fmt"

func rotate(nums []int, k int) {
	/* 1、切片法  时间负责度 O(n) 空间复杂度 O(n)
	if len(nums) <= 1 || k == 0 || k % len(nums) == 0{
		return
	}

	k %= len(nums)

	tmp := append(nums[len(nums) - k:], nums[:len(nums) - k]...)
	copy(nums, tmp)
	*/

	// 2、三次翻转法  时间负责度 O(n) 空间复杂度 O(1)
	n := len(nums)
	if n <= 1 || k == 0 {
		return
	}
	k %= n // 处理k大于数组长度的情况

	// 自定义翻转函数
	reverse := func(arr []int, start, end int) {
		for start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}

	reverse(nums, 0, n-1) // 整体翻转
	reverse(nums, 0, k-1) // 前k个元素翻转
	reverse(nums, k, n-1) // 剩余元素翻转
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	nums2 := []int{-1, -100, 3, 99}

	rotate(nums1, 3)
	rotate(nums2, 2)

	fmt.Println(nums1) // [5,6,7,1,2,3,4]
	fmt.Println(nums2) // [3,99,-1,-100]
}
