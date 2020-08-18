/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func nthSuperUglyNumber(n int, primes []int) int {
	if len(primes) == 0 || n == 1 {
		return 1
	}

	h := new(IntHeap)
	heap.Push(h, 1)
	n--
	for n > 0 {
		// fmt.Println(h)
		value := heap.Pop(h).(int)
		for (h.Len() > 0) && value == (*h)[0] {
			heap.Pop(h)
		}

		for i := 0; i < len(primes); i++ {
			heap.Push(h, value*primes[i])
		}
		n--
	}

	return heap.Pop(h).(int)
}

// @lc code=end
