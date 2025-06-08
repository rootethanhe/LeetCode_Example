package main

import "fmt"

// 方法1：逐位检查法
func hammingWeight1(n uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		mask := uint32(1 << i)
		if n&mask != 0 { // 检查第i位是否为1
			count++
		}
	}
	return count
}

// 方法2：清除最低位1法
func hammingWeight2(n uint32) int {
	count := 0
	for n != 0 {
		n = n & (n - 1) // 清除最低位的1
		count++
	}
	return count
}

// 方法3：SWAR算法（优化避免溢出）
func hammingWeight3(n uint32) int {
	// 步骤1：每2位分组统计（结果存储在每2位中）
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	// 步骤2：每4位分组统计（合并相邻2位组）
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	// 步骤3：每8位分组统计（合并相邻4位组）
	n = (n & 0x0F0F0F0F) + ((n >> 4) & 0x0F0F0F0F)
	// 步骤4：避免乘法溢出的合并
	n = (n & 0x00FF00FF) + ((n >> 8) & 0x00FF00FF) // 合并高/低16位
	n = (n & 0x0000FFFF) + (n >> 16)               // 最终合并
	return int(n)
}

// 方法4：查表法
var table = func() [256]int {
	table := [256]int{}
	for i := 0; i < 256; i++ {
		count := 0
		num := uint8(i)
		for num != 0 {
			num &= num - 1 // 清除最低位1
			count++
		}
		table[i] = count
	}
	return table
}()

func hammingWeight4(n uint32) int {
	return table[n&0xFF] + // 最低8位
		table[(n>>8)&0xFF] + // 次低8位
		table[(n>>16)&0xFF] + // 次高8位
		table[n>>24] // 最高8位
}

func main() {
	tests := []struct {
		name     string
		n        uint32
		expected int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 1},
		{"n=11", 11, 3},
		{"n=128", 128, 1},
		{"n=0xFFFFFFFF", 0xFFFFFFFF, 32}, // 全1
		{"n=2147483645", 2147483645, 30},
	}

	for _, test := range tests {
		fmt.Printf("输入: %-12s n=%d\n", test.name, test.n)
		fmt.Printf("逐位检查法: %d\n", hammingWeight1(test.n))
		fmt.Printf("清除最低位1法: %d\n", hammingWeight2(test.n))
		fmt.Printf("SWAR算法: %d\n", hammingWeight3(test.n))
		fmt.Printf("查表法: %d\n\n", hammingWeight4(test.n))
	}
}
