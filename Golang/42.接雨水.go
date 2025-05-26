package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0 // 边界处理：至少需要3根柱子才能形成凹槽
	}

	left, right := 0, n-1 // 双指针初始化
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		// 选择较矮的一侧处理(木桶效应)
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left] // 更新左侧最高挡板
			} else {
				water += leftMax - height[left] // 当前积水量
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right] // 更新右侧最高挡板
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 输出: 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 输出: 9
}
