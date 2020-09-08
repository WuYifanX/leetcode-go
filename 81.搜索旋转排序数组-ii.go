// package main

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (35.81%)
 * Likes:    209
 * Dislikes: 0
 * Total Accepted:    39K
 * Total Submissions: 108.8K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
 *
 * 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
 *
 * 示例 1:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 0
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 3
 * 输出: false
 *
 * 进阶:
 *
 *
 * 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
 * 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 *
 *
 */

//  0 0 0 1 1 1 1
//  0 1 1 1 1 0 0
// @lc code=start
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	return innerSearch(nums, 0, len(nums)-1, target)
}

func innerSearch(nums []int, start, end, target int) bool {
	left, right := start, end
	var middle int
	if left > right {
		return false
	}

	for left < right {
		if nums[left] < nums[right] {
			return binarySearch(nums, start, end, target)
		}

		middle = left + (right-left)/2

		return innerSearch(nums, start, middle, target) || innerSearch(nums, middle+1, end, target)
	}

	return nums[left] == target
}

func binarySearch(nums []int, start, end, target int) bool {
	left, right := start, end

	var middle int

	for left < right {

		middle = left + (right-left)/2

		if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return nums[left] == target
}

// @lc code=end

func main() {
	search([]int{2, 5, 6, 0, 0, 1, 2}, 0)
}
