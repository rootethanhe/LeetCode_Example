package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	var stack []string
	parts := strings.Split(path, "/") // 分割路径为多个部分

	for _, part := range parts {
		// 跳过空字符串和当前目录标识
		if part == "" || part == "." {
			continue
		}
		// 处理上级目录
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // 弹出栈顶元素(目录回退)
			}
		} else {
			stack = append(stack, part) // 有效目录名压入栈
		}
	}

	// 拼接最终路径
	result := "/" + strings.Join(stack, "/")
	// 处理根目录为空的情况(如输入 "/../")
	if len(result) == 0 || result[0] != '/' {
		return "/"
	}

	return result
}

func main() {
	// 测试用例
	testCases := []struct {
		input  string
		output string
	}{
		{"/home/", "/home"},
		{"/home//foo/", "/home/foo"},
		{"/../", "/"},
		{"/a/./b/../../c/", "/c"},
		{"/.../a/../b/c/../d/./", "/.../b/d"},
	}

	for _, tc := range testCases {
		res := simplifyPath(tc.input)
		fmt.Printf("输入: %-25s 期望输出: %-10s 实际输出: %-10s\n", tc.input, tc.output, res)
	}
}
