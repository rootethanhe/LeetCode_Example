package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 初始化数据结构
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	order := make([]int, 0, numCourses)

	// 构建图与入度统计
	for _, p := range prerequisites {
		ai, bi := p[0], p[1]
		graph[bi] = append(graph[bi], ai) // bi → ai
		indegree[ai]++
	}

	// 初始化队列（入度为0的课程）
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS拓扑排序
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		order = append(order, course)

		// 更新后续课程入度
		for _, neighbor := range graph[course] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 环检测
	if len(order) != numCourses {
		return []int{}
	}

	return order
}

func main() {
	// 示例1：简单依赖
	fmt.Println(findOrder(2, [][]int{{1, 0}})) // [0,1]

	// 示例2：多依赖路径
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	// 可能输出: [0,1,2,3] 或 [0,2,1,3]

	// 示例3：无依赖单课程
	fmt.Println(findOrder(1, [][]int{})) // [0]

	// 环测试：循环依赖
	fmt.Println(findOrder(2, [][]int{{1, 0}, {0, 1}})) // []

	// 复杂无环测试
	fmt.Println(findOrder(3, [][]int{{1, 0}, {2, 1}})) // [0,1,2]
}
