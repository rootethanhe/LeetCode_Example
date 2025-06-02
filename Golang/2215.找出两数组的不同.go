package main

import "fmt"

func findDifference(nums1, nums2 []int) [][]int {
	// 构建两个集合（自动去重）
	set1 := make(map[int]struct{})
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}

	set2 := make(map[int]struct{})
	for _, num := range nums2 {
		set2[num] = struct{}{}
	}

	// 计算 nums1 有但 nums2 没有的元素
	diff1 := []int{}
	for num := range set1 {
		if _, exists := set2[num]; !exists {
			diff1 = append(diff1, num)
		}
	}

	// 计算 nums2 有但 nums1 没有的元素
	diff2 := []int{}
	for num := range set2 {
		if _, exists := set1[num]; !exists {
			diff2 = append(diff2, num)
		}
	}

	return [][]int{diff1, diff2}
}

func main() {
	// 输入：nums1 = [1,2,3], nums2 = [2,4,6]  输出：[[1,3],[4,6]]
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))

	// 输入：nums1 = [1,2,3,3], nums2 = [1,1,2,2]  输出：[[3],[]]
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}
