// package main

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (38.53%)
 * Likes:    935
 * Dislikes: 0
 * Total Accepted:    168.3K
 * Total Submissions: 433.5K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 *
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 *
 * 你可以假设数组中不存在重复的元素。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 示例 1:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 *
 *
 * 示例 2:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	return innerSearch(nums, 0, len(nums)-1, target)
}

func innerSearch(nums []int, start, end, target int) int {
	left, right := start, end
	var middle int

	for left < right {
		if nums[left] < nums[right] {
			return binarySearch(nums, left, right, target)
		}
		middle = left + (right-left)/2
		// fmt.Println(left, right, middle)

		leftSearch := innerSearch(nums, start, middle, target)
		if leftSearch != -1 {
			return leftSearch
		}

		rightSearch := innerSearch(nums, middle+1, right, target)
		if rightSearch != -1 {
			return rightSearch
		}
		return -1
	}

	if nums[left] == target {
		return left
	} else {
		return -1
	}
}

func binarySearch(nums []int, start, end, target int) int {
	left, right := start, end
	var middle int

	for left < right {
		middle = left + (right-left)/2

		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}

	if nums[left] == target {
		return left
	} else {
		return -1
	}
}

// @lc code=end

func main() {
	search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
}
