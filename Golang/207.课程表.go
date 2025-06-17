package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 初始化邻接表和入度数组
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	// 构建图和入度统计
	for _, p := range prerequisites {
		course := p[0]
		prereq := p[1]
		graph[prereq] = append(graph[prereq], course) // prereq -> course
		indegree[course]++                            // course的入度+1
	}

	// 初始化队列(入度为0的课程)
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS 拓扑排序
	count := 0 // 已处理课程计数
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count++

		// 更新后续节点入度
		for _, neighbor := range graph[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return count == numCourses
}

func main() {
	// 示例1：无环
	fmt.Println(canFinish(2, [][]int{{1, 0}})) // true

	// 示例2：有环
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}})) // false

	// 复杂无环
	fmt.Println(canFinish(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})) // true

	// 复杂有环
	fmt.Println(canFinish(3, [][]int{{1, 0}, {2, 1}, {0, 2}})) // false
}
