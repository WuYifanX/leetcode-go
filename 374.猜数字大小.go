package main

import "math"

func guess(a int) int {
	return 0
}

/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	right := math.MaxInt32
	left := 0

	for left < right {
		middle := left + (right-left+1)/2
		if guess(middle) == -1 {
			right = middle - 1
		} else {
			left = middle
		}
	}
	return left
}

// @lc code=end
