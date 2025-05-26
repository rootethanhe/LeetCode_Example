package main

import "fmt"

func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(words) == 0 {
		return res
	}
	wordLen := len(words[0])
	numWords := len(words)
	totalLen := wordLen * numWords
	sLen := len(s)
	if sLen < totalLen {
		return res
	}

	// 统计单词出现次数
	wordCount := make(map[string]int)
	for _, w := range words {
		wordCount[w]++
	}

	// 遍历每个起始偏移（0 到 wordLen-1）
	for i := 0; i < wordLen; i++ {
		left, right := i, i
		currentCount := make(map[string]int)
		valid := 0 // 记录满足次数要求的单词种类数

		for right+wordLen <= sLen {
			// 提取当前单词并移动右指针
			currentWord := s[right : right+wordLen]
			right += wordLen

			// 如果单词在目标列表中
			if targetCount, exists := wordCount[currentWord]; exists {
				currentCount[currentWord]++
				// 当前单词次数刚好满足要求时，valid 增加
				if currentCount[currentWord] == targetCount {
					valid++
				}

				// 若当前单词次数超过要求，移动左指针直到次数合理
				for currentCount[currentWord] > targetCount {
					leftWord := s[left : left+wordLen]
					left += wordLen
					if currentCount[leftWord] == wordCount[leftWord] {
						valid--
					}
					currentCount[leftWord]--
				}

				// 当窗口长度等于 totalLen 时检查条件
				if right-left == totalLen {
					// 有效单词种类数等于哈希表键数时，记录结果
					if valid == len(wordCount) {
						res = append(res, left)
					}
					// 移动左指针，移除最左侧单词
					leftWord := s[left : left+wordLen]
					if currentCount[leftWord] == wordCount[leftWord] {
						valid--
					}
					currentCount[leftWord]--
					left += wordLen
				}
			} else {
				// 遇到无效单词，重置窗口
				currentCount = make(map[string]int)
				valid = 0
				left = right
			}
		}
	}

	return res
}

func main() {
	// 示例 1
	s1 := "barfoothefoobarman"
	words1 := []string{"foo", "bar"}
	fmt.Println(findSubstring(s1, words1)) // 输出: [0 9]

	// 示例 2
	s2 := "wordgoodgoodgoodbestword"
	words2 := []string{"word", "good", "best", "word"}
	fmt.Println(findSubstring(s2, words2)) // 输出: []

	// 示例 3
	s3 := "barfoofoobarthefoobarman"
	words3 := []string{"bar", "foo", "the"}
	fmt.Println(findSubstring(s3, words3)) // 输出: [6 9 12]
}
