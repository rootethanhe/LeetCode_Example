package main

import "fmt"

func trailingZeroes(n int) int {
	count := 0
	// 循环累加5的幂次贡献的因子数
	for n >= 5 {
		n /= 5     // 更新n为当前5的倍数个数
		count += n // 累加当前层级的因子5数量
	}
	return count
}

func main() {
	tests := []struct {
		input  int
		expect int
	}{
		{3, 0},    // 3! = 6 → 无尾随零
		{5, 1},    // 5! = 120 → 1个尾随零
		{0, 0},    // 0! = 1 → 无尾随零
		{25, 6},   // 25! → 6个尾随零（含25额外贡献）
		{100, 24}, // 100! → 24个尾随零
	}

	for _, test := range tests {
		res := trailingZeroes(test.input)
		fmt.Printf("n=%d → %d (expected %d)\n", test.input, res, test.expect)
	}
}
