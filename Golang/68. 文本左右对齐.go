package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	start := 0
	n := len(words)

	for start < n {
		// 确定当前行的结束索引和单词总长度
		end, sumWordLen := findEnd(words, start, maxWidth)
		// 生成当前行的字符串
		line := buildLine(words, start, end, sumWordLen, maxWidth, end == n-1)
		result = append(result, line)
		start = end + 1
	}
	return result
}

// 找到当前行的结束索引和单词总长度
func findEnd(words []string, start, maxWidth int) (int, int) {
	sumWordLen := 0
	end := start
	for end < len(words) {
		currentWordLen := sumWordLen + len(words[end]) // 尝试添加当前单词
		requiredSpace := end - start                   // 当前行已包含的空格数（单词数-1）
		totalLength := currentWordLen + requiredSpace  // 总长度（单词+空格）
		if totalLength > maxWidth {                    // 超过最大宽度则停止
			return end - 1, sumWordLen
		}
		sumWordLen = currentWordLen // 更新总单词长度
		end++
	}
	return end - 1, sumWordLen // 处理到最后一个单词的情况
}

// 构建单行字符串
func buildLine(words []string, start, end, sumWordLen, maxWidth int, isLastLine bool) string {
	var builder strings.Builder
	wordCount := end - start + 1
	totalSpaces := maxWidth - sumWordLen

	// 处理最后一行或单单词行：左对齐，末尾补空格
	if wordCount == 1 || isLastLine {
		for i := start; i <= end; i++ {
			if i > start {
				builder.WriteByte(' ')
			}
			builder.WriteString(words[i])
		}
		remainingSpaces := maxWidth - builder.Len()
		builder.WriteString(strings.Repeat(" ", remainingSpaces))
		return builder.String()
	}

	// 普通行：均匀分配空格，左侧优先多放
	slots := wordCount - 1
	baseSpace := totalSpaces / slots
	extraSpace := totalSpaces % slots

	builder.WriteString(words[start])
	for i := 0; i < slots; i++ {
		space := baseSpace
		if i < extraSpace { // 前 extraSpace 个间隙多1空格
			space++
		}
		builder.WriteString(strings.Repeat(" ", space))
		builder.WriteString(words[start+i+1])
	}
	return builder.String()
}

func main() {
	// 示例1
	words1 := []string{"This", "is", "an", "example", "of", "text", "justification."}
	fmt.Println(fullJustify(words1, 16))
	// 输出: [This    is    an example  of text justification.  ]

	// 示例2
	words2 := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	fmt.Println(fullJustify(words2, 16))
	// 输出: [What   must   be acknowledgment   shall be        ]
}
