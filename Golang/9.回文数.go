package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	// 处理特殊情况：负数或末位为0的非零数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0
	// 反转后半部分，直到原数小于等于反转后的数
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// 偶数位直接比较，奇数位需去掉中间位
	return x == reversed || x == reversed/10
}

func main() {
	fmt.Println(isPalindrome(121))   // true
	fmt.Println(isPalindrome(-121))  // false
	fmt.Println(isPalindrome(10))    // false
	fmt.Println(isPalindrome(12321)) // true
}
