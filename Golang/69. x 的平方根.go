package main

import "fmt"

func mySqrt(x int) int {
	if x <= 1 {
		return x // 直接处理 0 和 1 的边界情况
	}
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2 // 避免整数溢出
		square := mid * mid
		if square == x {
			return mid
		} else if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right // 最终right是最大整数解
}

func main() {
	fmt.Println(mySqrt(4)) // 2
	fmt.Println(mySqrt(8)) // 2
}
