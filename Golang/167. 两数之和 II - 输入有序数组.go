package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1 // 初始化双指针
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target { // 找到唯一解
			return []int{left + 1, right + 1} // 下标从1开始，需+1
		} else if sum < target { // 和过小，左指针右移
			left++
		} else { // 和过大，右指针左移
			right--
		}
	}
	return []int{-1, -1} // 实际不会执行（题目保证存在解）
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [1,2]
	fmt.Println(twoSum([]int{2, 3, 4}, 6))      // [1,3]
	fmt.Println(twoSum([]int{-1, 0}, -1))       // [1, 2]
}
