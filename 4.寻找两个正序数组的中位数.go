// package main

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (37.10%)
 * Likes:    2297
 * Dislikes: 0
 * Total Accepted:    158.4K
 * Total Submissions: 427K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 *
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	left := (len1 + len2 + 1) / 2
	right := (len1 + len2 + 2) / 2
	leftValue := findKth(nums1, 0, nums2, 0, left)
	rightValue := findKth(nums1, 0, nums2, 0, right)

	return float64(leftValue+rightValue) / 2
}

func findKth(nums1 []int, startIndex1 int, nums2 []int, startIndex2 int, k int) int {
	// 如果任意一个数组已经删除完毕了, 然后剩下的第k个就直接是了
	if startIndex1 >= len(nums1) {
		return nums2[startIndex2+k-1]
	} else if startIndex2 >= len(nums2) {
		return nums1[startIndex1+k-1]
	}

	if k == 1 {
		return minOf(nums1[startIndex1], nums2[startIndex2])
	}

	len1 := len(nums1) - startIndex1
	i := startIndex1 + minOf(len1, k/2) - 1
	len2 := len(nums2) - startIndex2
	j := startIndex2 + minOf(len2, k/2) - 1

	if nums1[i] > nums2[j] {
		return findKth(nums1, startIndex1, nums2, j+1, k-minOf(len2, k/2))
	} else {
		return findKth(nums1, i+1, nums2, startIndex2, k-minOf(len1, k/2))
	}
}

func minOf(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
// func main() {
// 	fmt.Println(findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6}))
// }
