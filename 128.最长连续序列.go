// package main

// import "fmt"

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 *
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (51.66%)
 * Likes:    483
 * Dislikes: 0
 * Total Accepted:    71.1K
 * Total Submissions: 137.6K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组，找出最长连续序列的长度。
 *
 * 要求算法的时间复杂度为 O(n)。
 *
 * 示例:
 *
 * 输入: [100, 4, 200, 1, 3, 2]
 * 输出: 4
 * 解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	maps := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		maps[nums[i]] = i
	}

	result := 0
	currentResult := 0
	fmt.Println(maps)
	for k, _ := range maps {
		// find Head key
		currentKey := k

		for {
			_, exist := maps[currentKey]

			if exist {
				currentKey = currentKey - 1
			} else {
				break
			}
		}

		currentKey++
		currentResult = 0
		for {
			_, exist := maps[currentKey]

			if exist {
				delete(maps, currentKey)
				currentKey++
				currentResult++
			} else {
				result = maxOf(result, currentResult)
				break
			}
		}

	}

	return result
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
// func main() {
// 	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
// }
