package main

import "fmt"

func isHappy(n int) bool {
	/* **** 1、哈希集合法 ****
	nMap := make(map[int]int)
	for {
		// 如果存在重复的值，则是无限循环的
		n = bitSquareSum(n)
		if n == 1 {
			// 值为1时，是快乐数
			return true
		} else if _, exists := nMap[n]; exists {
			return false
		} else {
			nMap[n] = 1
		}
	}
	*/

	// **** 2、快慢指针法 ****
	// 边界条件：直接返回1的情况
	if n == 1 {
		return true
	}

	// 快慢指针初始化
	slow, fast := n, bitSquareSum(n)

	// 循环检测（快指针可能先到1）
	for fast != 1 && slow != fast {
		slow = bitSquareSum(slow)               // 慢指针走一步
		fast = bitSquareSum(bitSquareSum(fast)) // 快指针走两步
	}

	return fast == 1
}

// 计算数字的各位平方和
func bitSquareSum(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10 // 取各位数
		sum += digit * digit
		n /= 10 // 移除已处理的位数
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19)) // true
	fmt.Println(isHappy(2))  // false
}
