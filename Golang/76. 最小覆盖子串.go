package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// 统计 t 的字符频率
	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	windowFreq := make(map[byte]int) // 窗口内字符频率
	left, right := 0, 0              // 滑动窗口左右指针
	valid := 0                       // 满足 t 频率要求的字符种类数
	minLen := math.MaxInt32          // 最小窗口长度
	start := 0                       // 最小窗口起始位置

	for right < len(s) {
		c := s[right]
		right++

		// 若当前字符在 t 中，更新窗口统计
		if _, exists := tFreq[c]; exists {
			windowFreq[c]++
			if windowFreq[c] == tFreq[c] {
				valid++
			}
		}

		// 当窗口包含所有 t 的字符时，尝试收缩左边界
		for valid == len(tFreq) {
			// 更新最小窗口
			if right-left < minLen {
				minLen = right - left
				start = left
			}

			// 移动左指针
			d := s[left]
			left++
			if _, exists := tFreq[d]; exists {
				if windowFreq[d] == tFreq[d] {
					valid-- // 移除后不再满足条件
				}
				windowFreq[d]--
			}
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}

func main() {
	// 示例 1
	s1 := "ADOBECODEBANC"
	t1 := "ABC"
	fmt.Println(minWindow(s1, t1)) // 输出: "BANC"

	// 示例 2
	s2 := "a"
	t2 := "a"
	fmt.Println(minWindow(s2, t2)) // 输出: "a"

	// 示例 3
	s3 := "a"
	t3 := "aa"
	fmt.Println(minWindow(s3, t3)) // 输出: ""
}
