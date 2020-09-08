
/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (40.17%)
 * Likes:    578
 * Dislikes: 0
 * Total Accepted:    129.7K
 * Total Submissions: 321.8K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 如果数组中不存在目标值，返回 [-1, -1]。
 *
 * 示例 1:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 *
 * 示例 2:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 *
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstPosition := binarySearchFirstTarget(nums, target)

	if nums[firstPosition] != target {
		return []int{-1, -1}
	}

	var secondPosition int
	if nums[len(nums)-1] == target {
		secondPosition = len(nums) - 1
	} else {
		secondPosition = binarySearchFirstTarget(nums, target+1) - 1
	}

	return []int{firstPosition, secondPosition}
}

func binarySearchFirstTarget(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var middle int

	for left < right {
		middle = left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left

}

// @lc code=end
