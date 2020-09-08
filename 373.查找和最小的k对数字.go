// package main

import (
	"container/heap"
)

/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的K对数字
 */

// @lc code=start

type Pair struct {
	x int
	y int
}

type PairHeap []Pair

func (h PairHeap) Len() int { return len(h) }

func (h PairHeap) Less(i, j int) bool {
	return h[i].x+h[i].y < h[j].x+h[j].y
}
func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, 0)

	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return result
	}

	h := new(PairHeap)
	heap.Init(h)
	pushAllResultIntoHeap(h, nums1, nums2)

	for len(result) < k && h.Len() != 0 {
		current := heap.Pop(h).(Pair)
		result = append(result, []int{current.x, current.y})
	}
	return result
}

func pushAllResultIntoHeap(h *PairHeap, nums1, nums2 []int) {
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			heap.Push(h, Pair{nums1[i], nums2[j]})
		}
	}
}

// @lc code=end
