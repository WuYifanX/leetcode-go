package main

import (
	"container/list"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	graph := buildGraphList(n, edges)
	isVisited := buildGraphMatrix(len(*graph), [][]int{})

	heightMap := make([]int, n)
	minHeightForCurrent := n

	for i := 0; i < n; i++ {
		heightMap[i] = minHeight(i, graph, isVisited, minHeightForCurrent)
		if heightMap[i] < minHeightForCurrent {
			minHeightForCurrent = heightMap[i]
		}
	}
	// fmt.Println(heightMap)

	// for i := 0; i < n; i++ {
	// 	if minHeightForCurrent >= heightMap[i] {
	// 		minHeightForCurrent = heightMap[i]
	// 	}
	// }
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if minHeightForCurrent == heightMap[i] {
			result = append(result, i)
		}
	}

	return result

}

func minHeight(v int, graph *map[int][]int, isVisited *[][]int, minHeightForCurrent int) int {

	queue := list.New()

	queue.PushBack(v)
	height := 0
	size := 0
	for queue.Len() != 0 {
		size = queue.Len()
		height++

		if height > minHeightForCurrent {
			return height
		}

		for i := 0; i < size; i++ {
			element := queue.Front()
			current := element.Value.(int)
			queue.Remove(element)
			for _, next := range (*graph)[current] {
				if (*isVisited)[current][next] != v {
					queue.PushBack(next)
					(*isVisited)[current][next] = v
					(*isVisited)[next][current] = v
				}
			}

		}
	}
	return height
}

func buildGraphList(vertex int, edges [][]int) *map[int][]int {
	graph := make(map[int][]int)

	for v := 0; v < vertex; v++ {
		graph[v] = make([]int, 0)
	}

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return &graph
}

func buildGraphMatrix(vertex int, edges [][]int) *[][]int {
	graph := make([][]int, vertex)

	for v := 0; v < vertex; v++ {
		graph[v] = make([]int, vertex)
	}

	for i := 0; i < vertex; i++ {
		for j := 0; j < vertex; j++ {
			graph[i][j] = -1
		}
	}
	return &graph
}

// @lc code=end

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{0, 1}, {1, 2}, {1, 3}}))
	// fmt.Println(findMinHeightTrees(6, [][]int{{0, 1}, {2, 0}, {0, 3}, {3, 4}, {4, 5}}))
}
