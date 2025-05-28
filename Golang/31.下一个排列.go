package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)

	// 步骤1：从右向左找第一个断点i(nums[i]<nums[i+1])
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 步骤2：找到交换目标j
	if i >= 0 {
		j := n - 1
		for j > i && nums[j] <= nums[i] { // 找比 nums[i] 大的最右侧元素
			j--
		}
		nums[i], nums[j] = nums[j], nums[i] // 交换
	}

	// 步骤3：反转 i+1 之后的降序为升序
	reverse(nums, i+1, n-1)
}

// 反转数组的指定区间
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	// 输入：nums = [1,2,3]  输出：[1,3,2]
	nums1 := []int{1, 2, 3}
	nextPermutation(nums1)
	fmt.Println(nums1)

	// 输入：nums = [3,2,1]  输出：[1,2,3]
	nums2 := []int{3, 2, 1}
	nextPermutation(nums2)
	fmt.Println(nums2)

	// 输入：nums = [1,1,5]  输出：[1,5,1]
	nums3 := []int{1, 1, 5}
	nextPermutation(nums3)
	fmt.Println(nums3)

}
