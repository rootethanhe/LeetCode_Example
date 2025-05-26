package main

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}
	backtrack(&result, "", 0, 0, n)

	return result
}

// backtrack 回溯函数
// result: 结果数组的指针
// current: 当前构建的字符串
// open: 已使用的左括号数
// close: 已使用的右括号数
// max: 目标括号对数 n
func backtrack(result *[]string, current string, open, close, max int) {
	// 终止条件: 当前字符串长度到达2n
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	// 优先添加左括号(若未达到上限)
	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}

	// 在右括号不足左括号时，添加右括号
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}
}

func main() {
	// 输入：n = 3
	fmt.Println(generateParenthesis(3))
	// 输出：["((()))","(()())","(())()","()(())","()()()"]

	// 输入：n = 1
	fmt.Println(generateParenthesis(1))
	// ["()"]
}
