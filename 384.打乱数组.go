package main

import (
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start
type Solution struct {
	origin []int
}

func Constructor(nums []int) Solution {
	solution := Solution{nums}
	return solution
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origin
}

func swapInNums(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func randomRange(start, end int) int {
	return rand.Intn(end-start+1) + start
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {

	result := make([]int, len(this.origin))

	copy(result, this.origin)

	for i := 0; i < len(result); i++ {
		swapInNums(&result, i, randomRange(i, len(result)-1))
	}
	return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end
