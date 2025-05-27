package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// 从指定目录中随机获取一个文件名
func getRandomFile(dirPath string) (string, error) {
	// 读取目录内容
	entries, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return "", fmt.Errorf("read dir(%s) fail", dirPath)
	}

	// 过滤文件
	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, strings.TrimSuffix(entry.Name(), ".go"))
		}
	}

	// 处理无文件情况
	if len(files) == 0 {
		return "", fmt.Errorf("no files found in dir(%s)", dirPath)
	}

	// 设置随机种子(避免重复性)
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(files))

	return files[randomIndex], nil
}

func main() {
	// 随机抽取一个练习题
	dirPath := "./Golang"
	filename, err := getRandomFile(dirPath)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(filename)
}
