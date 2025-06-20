package main

import "fmt"

func singleNumberStateMachine(nums []int) int {
	var one, two int
	for _, num := range nums {
		two |= one & num   // 同步 one 到 two（出现两次）
		one ^= num         // 更新出现奇数次的位
		three := one & two // 标记出现三次的位
		one &^= three      // 重置 one（清除三次出现的位）
		two &^= three      // 重置 two
	}
	return one
}

func singleNumberBitCount(nums []int) int {
	var ans int
	for i := 0; i < 32; i++ {
		cnt := 0
		for _, num := range nums {
			cnt += (num >> i) & 1 // 统计第 i 位 1 的总数
		}
		if cnt%3 != 0 { // 目标数字在此位为 1
			ans |= 1 << i // 设置结果的第 i 位
		}
	}
	return ans
}

func main() {
	nums1 := []int{2, 2, 3, 2}
	nums2 := []int{0, 1, 0, 1, 0, 1, 99}

	fmt.Println("状态机解法:")
	fmt.Println(singleNumberStateMachine(nums1)) // 输出 3
	fmt.Println(singleNumberStateMachine(nums2)) // 输出 99

	fmt.Println("逐位统计解法:")
	fmt.Println(singleNumberBitCount(nums1)) // 输出 3
	fmt.Println(singleNumberBitCount(nums2)) // 输出 99
}
