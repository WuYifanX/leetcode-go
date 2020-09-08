// package main

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (51.54%)
 * Likes:    248
 * Dislikes: 0
 * Total Accepted:    76.5K
 * Total Submissions: 148.3K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
 *
 * 请找出其中最小的元素。
 *
 * 你可以假设数组中不存在重复元素。
 *
 * 示例 1:
 *
 * 输入: [3,4,5,1,2]
 * 输出: 1
 *
 * 示例 2:
 *
 * 输入: [4,5,6,7,0,1,2]
 * 输出: 0
 *
 */

// @lc code=start
func findMin(nums []int) int {
	return innerFindMin(nums, 0, len(nums)-1)
}
func innerFindMin(nums []int, start, end int) int {
	left, right := start, end

	var middle int
	if left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		middle = left + (right-left)/2

		return minOf(innerFindMin(nums, start, middle), innerFindMin(nums, middle+1, end))

	} else {
		return nums[left]
	}
}

func minOf(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
