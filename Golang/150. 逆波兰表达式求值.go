package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	operators := map[string]bool{"+": true, "-": true, "*": true, "/": true}

	for _, token := range tokens {
		if _, ok := operators[token]; ok {
			// 弹出右操作数和左操作数
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// 执行运算
			var res int
			switch token {
			case "+":
				res = left + right
			case "-":
				res = left - right
			case "*":
				res = left * right
			case "/":
				res = left / right
			}
			stack = append(stack, res)
		} else {
			// 数字入栈
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))                                             // 输出 9
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))                                            // 输出 6
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // 输出 22
}
