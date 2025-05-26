package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		// 计算当前水量(木桶效应，取较矮的柱子)
		currentHeight := minHeight(height[left], height[right])
		currentWidth := right - left
		currentWater := currentHeight * currentWidth

		// 更新最大值
		if currentWater > maxWater {
			maxWater = currentWater
		}

		// 移动较矮的指针(贪心算法)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func minHeight(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                      // 1
}
