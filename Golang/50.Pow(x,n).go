package main

import "fmt"

func myPow(x float64, n int) float64 {
	// 边界: x为0时直接返回0
	if x == 0 {
		return 0
	}

	// 边界: n为0时返回1
	if n == 0 {
		return 1
	}

	// 处理负指数: 用int64避免溢出
	longN := int64(n)
	if longN < 0 {
		x = 1 / x
		longN = -longN
	}

	res := 1.0
	// 快速幂迭代
	for longN > 0 {
		if longN&1 == 1 { // 当前二进制为1
			res *= x
		}
		x *= x      // 底数平方
		longN >>= 1 // 指数右移
	}

	return res
}

func main() {
	// 输入：x = 2.00000, n = 10  输出：1024.00000
	fmt.Println(myPow(2.00000, 10))

	// 输入：x = 2.10000, n = 3  输出：9.26100
	fmt.Println(myPow(2.10000, 3))

	// 输入：x = 2.00000, n = -2  输出：0.25000
	// 解释：2-2 = 1/22 = 1/4 = 0.25
	fmt.Println(myPow(2.00000, -2))
}
