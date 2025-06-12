package main

import "fmt"

// 方法1：逐位循环法
func reverseBitsLoop(n uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = (res << 1) | (n & 1) // 左移结果并添加最低位
		n >>= 1                    // 移除已处理的最低位
	}
	return res
}

// 方法2：分治交换法
func reverseBitsDivide(n uint32) uint32 {
	// 交换32位中的高16位和低16位
	n = (n >> 16) | (n << 16)
	// 交换每16位中的高8位和低8位
	n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8)
	// 交换每8位中的高4位和低4位
	n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4)
	// 交换每4位中的高2位和低2位
	n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2)
	// 交换每2位中的高1位和低1位
	n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1)
	return n
}

func main() {
	testCases := []struct {
		input  uint32
		expect uint32
	}{
		{0b00000010100101000001111010011100, 964176192},  // 示例1
		{0b11111111111111111111111111111101, 3221225471}, // 示例2
		{0b10000000000000000000000000000000, 1},          // 边界测试1
		{0xFFFFFFFF, 0xFFFFFFFF},                         // 全1测试
	}

	fmt.Println("测试结果:")
	for _, tc := range testCases {
		res1 := reverseBitsLoop(tc.input)
		res2 := reverseBitsDivide(tc.input)
		fmt.Printf("输入: %032b\n", tc.input)
		fmt.Printf("逐位法: %032b (%d)\n", res1, res1)
		fmt.Printf("分治法: %032b (%d)\n", res2, res2)
		fmt.Printf("预期值: %d\n", tc.expect)
		fmt.Println("匹配:", res1 == tc.expect && res2 == tc.expect, "\n")
	}
}
