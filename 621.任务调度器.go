// package main

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 *
 * https://leetcode-cn.com/problems/task-scheduler/description/
 *
 * algorithms
 * Medium (50.51%)
 * Likes:    349
 * Dislikes: 0
 * Total Accepted:    29.4K
 * Total Submissions: 58.1K
 * Testcase Example:  '["A","A","A","B","B","B"]\n2'
 *
 * 给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26
 * 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU
 * 在任何一个单位时间内都可以执行一个任务，或者在待命状态。
 *
 * 然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
 *
 * 你需要计算完成所有任务所需要的最短时间。
 *
 *
 *
 * 示例 ：
 *
 * 输入：tasks = ["A","A","A","B","B","B"], n = 2
 * 输出：8
 * 解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
 * ⁠    在本示例中，两个相同类型任务之间必须间隔长度为 n = 2
 * 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。
 *
 *
 *
 * 提示：
 *
 *
 * 任务的总个数为 [1, 10000]。
 * n 的取值范围为 [0, 100]。
 *
 *
 */

// @lc code=start

// An IntHeap is a min-heap of ints.

type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	if len(tasks) <= 1 {
		return len(tasks)
	}

	taskMaps := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		taskMaps[tasks[i]-'A'] += 1
	}

	h := new(IntHeap)
	for i := 0; i < 26; i++ {
		if taskMaps[i] != 0 {
			heap.Push(h, []int{i, taskMaps[i]})
		}
	}

	recentUsedMaps := make([]int, 26)

	for i := 0; i < 26; i++ {
		recentUsedMaps[i] = -1
	}
	result := 0
	tempResult := 0
	temp := make([]int, 26)
	// fmt.Println(taskMaps)
	for h.Len() != 0 {
		// fmt.Println(h)
		tempResult = result
		for h.Len() != 0 {
			topElement := heap.Pop(h)
			topValue := topElement.([]int)
			// fmt.Println(topValue)
			if recentUsedMaps[topValue[0]] == -1 || result-recentUsedMaps[topValue[0]] >= n {
				result++
				topValue[1] -= 1
				if topValue[1] > 0 {
					temp[topValue[0]] = topValue[1]
				}
				recentUsedMaps[topValue[0]] = result
				break
			} else {
				temp[topValue[0]] = topValue[1]
			}
		}

		// idle
		if tempResult == result {
			result++
		}
		// fmt.Println(result, recentUsedMaps, temp, h.Len())

		for i := 0; i < 26; i++ {
			if temp[i] > 0 {
				heap.Push(h, []int{i, temp[i]})
				temp[i] = 0
			}
		}

	}

	return result
}

// @lc code=end

func main() {

	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}
