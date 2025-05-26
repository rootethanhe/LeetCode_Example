package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 处理奇数长度回文(如"aba")
		l1, r1 := expand(s, i, i)
		// 处理偶数长度回文(如"abba")
		l2, r2 := expand(s, i, i+1)
		// 更新最长回文边界
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

// 从中心向两边扩展，返回回文的最左/右边界
func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// 回退到最后一个满足条件的边界
	return left + 1, right - 1
}

func main() {
	// 输入：s = "babad"  输出："bab" / "aba" 也可
	fmt.Println(longestPalindrome("babad"))

	// 输入：s = "cbbd"  输出："bb"
	fmt.Println(longestPalindrome("cbbd"))
}
