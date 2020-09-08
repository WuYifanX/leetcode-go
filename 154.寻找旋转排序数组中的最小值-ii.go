package main

/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Hard (50.19%)
 * Likes:    173
 * Dislikes: 0
 * Total Accepted:    36.2K
 * Total Submissions: 72.2K
 * Testcase Example:  '[1,3,5]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
 *
 * 请找出其中最小的元素。
 *
 * 注意数组中可能存在重复的元素。
 *
 * 示例 1：
 *
 * 输入: [1,3,5]
 * 输出: 1
 *
 * 示例 2：
 *
 * 输入: [2,2,2,0,1]
 * 输出: 0
 *
 * 说明：
 *
 *
 * 这道题是 寻找旋转排序数组中的最小值 的延伸题目。
 * 允许重复会影响算法的时间复杂度吗？会如何影响，为什么？
 *
 *
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return innerFindMin(nums, 0, len(nums)-1)
}

func innerFindMin(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}

	if nums[start] < nums[end] {
		return nums[start]
	}

	middle := start + (end-start)/2

	return minOf(innerFindMin(nums, start, middle), innerFindMin(nums, middle+1, end))
}

func minOf(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
