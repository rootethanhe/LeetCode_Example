package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)

	ClosestSum := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		// 剪枝，跳过重复的nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			// 如果当前和等于目标值，直接返回
			if sum == target {
				return sum
			}
			// 更新最接近的和
			if abs(sum-target) < abs(ClosestSum-target) {
				ClosestSum = sum
			}

			// 根据和大小，移动指针
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return ClosestSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// 输入：nums = [-1,2,1,-4], target = 1  输出：2
	// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))

	// 输入：nums = [0,0,0], target = 1  输出：0
	// 解释：与 target 最接近的和是 0（0 + 0 + 0 = 0）
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}
