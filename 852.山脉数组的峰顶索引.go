/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 *
 * https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/description/
 *
 * algorithms
 * Easy (70.81%)
 * Likes:    109
 * Dislikes: 0
 * Total Accepted:    29.8K
 * Total Submissions: 42K
 * Testcase Example:  '[0,1,0]'
 *
 * 我们把符合下列属性的数组 A 称作山脉：
 *
 *
 * A.length >= 3
 * 存在 0 < i < A.length - 1 使得A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... >
 * A[A.length - 1]
 *
 *
 * 给定一个确定为山脉的数组，返回任何满足 A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... >
 * A[A.length - 1] 的 i 的值。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[0,1,0]
 * 输出：1
 *
 *
 * 示例 2：
 *
 * 输入：[0,2,1,0]
 * 输出：1
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= A.length <= 10000
 * 0 <= A[i] <= 10^6
 * A 是如上定义的山脉
 *
 *
 *
 *
 */

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	if len(arr) < 3 {
		return -1
	}

	left, right := 0, len(arr)-1
	var middle int
	for left < right-1 {
		middle = left + (right-left)/2

		// 1 2 1
		if arr[middle] > arr[middle-1] && arr[middle] > arr[middle+1] {
			return middle
		}

		// 1 2 3
		if arr[middle-1] < arr[middle] && arr[middle] < arr[middle+1] {
			left = middle + 1
		}

		// 3 2 1
		if arr[middle-1] > arr[middle] && arr[middle] > arr[middle+1] {
			right = middle - 1
		}
	}

	if left == right {
		return left
	} else {
		if arr[left] > arr[right] {
			return left
		} else {
			return right
		}
	}
}

// @lc code=end
