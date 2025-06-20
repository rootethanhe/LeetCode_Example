package main

import "fmt"

func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	// 右移直到 left 和 right 相等（找到公共前缀）
	for left < right {
		left >>= 1  // 去掉最低位
		right >>= 1 // 去掉最低位
		shift++     // 记录移位次数
	}
	// 左移补零，恢复公共前缀后的低位0
	return left << shift
}

func main() {
	tests := []struct {
		left, right int
		want        int
	}{
		{5, 7, 4},          // 示例1：二进制 101 & 111 → 100 (4)
		{0, 0, 0},          // 示例2：0 & 0 → 0
		{1, 2147483647, 0}, // 示例3：1(1) 和 2147483647(31个1) → 公共前缀0
		{10, 15, 8},        // 10(1010) & 15(1111) → 公共前缀1 → 1000(8)
	}

	for _, tt := range tests {
		res := rangeBitwiseAnd(tt.left, tt.right)
		fmt.Printf("[%d, %d] → %d (expected %d)\n", tt.left, tt.right, res, tt.want)
	}
}
