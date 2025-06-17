package main

import "fmt"

type Node struct {
	name    string
	product float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 1. 构建图
	graph := make(map[string]map[string]float64)
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]       // a->b权重
		graph[b][a] = 1.0 / values[i] // b->a权重
	}

	// 2. 处理查询
	results := make([]float64, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]

		// 变量不存在
		if graph[start] == nil || graph[end] == nil {
			results[i] = -1.0
			continue
		}

		// 相同变量
		if start == end {
			results[i] = 1.0
			continue
		}

		// 3. BFS搜索
		visited := make(map[string]bool)
		queue := []Node{{start, 1.0}}
		visited[start] = true
		found := false

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			// 遍历邻居节点
			for neighbor, weight := range graph[curr.name] {
				if visited[neighbor] {
					continue
				}

				newProduct := curr.product * weight

				// 找到目标
				if neighbor == end {
					results[i] = newProduct
					found = true
					break
				}

				// 加入队列继续搜索
				visited[neighbor] = true
				queue = append(queue, Node{neighbor, newProduct})
			}
			if found {
				break
			}
		}

		// 未找到路径
		if !found {
			results[i] = -1.0
		}
	}

	return results
}

func main() {
	// 示例1
	eq1 := [][]string{{"a", "b"}, {"b", "c"}}
	val1 := []float64{2.0, 3.0}
	queries1 := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(calcEquation(eq1, val1, queries1)) // [6, 0.5, -1, 1, -1]

	// 示例2
	eq2 := [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	val2 := []float64{1.5, 2.5, 5.0}
	queries2 := [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}
	fmt.Println(calcEquation(eq2, val2, queries2)) // [3.75, 0.4, 5, 0.2]

	// 示例3
	eq3 := [][]string{{"a", "b"}}
	val3 := []float64{0.5}
	queries3 := [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}
	fmt.Println(calcEquation(eq3, val3, queries3)) // [0.5, 2, -1, -1]
}
