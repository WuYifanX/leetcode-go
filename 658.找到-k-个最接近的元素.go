import (
	"container/list"
)

/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 *
 * https://leetcode-cn.com/problems/find-k-closest-elements/description/
 *
 * algorithms
 * Medium (43.93%)
 * Likes:    139
 * Dislikes: 0
 * Total Accepted:    13.5K
 * Total Submissions: 30.7K
 * Testcase Example:  '[1,2,3,4,5]\n4\n3'
 *
 * 给定一个排序好的数组，两个整数 k 和 x，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。如果有两个数与 x
 * 的差值一样，优先选择数值较小的那个数。
 *
 * 示例 1:
 *
 *
 * 输入: [1,2,3,4,5], k=4, x=3
 * 输出: [1,2,3,4]
 *
 *
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1,2,3,4,5], k=4, x=-1
 * 输出: [1,2,3,4]
 *
 *
 *
 *
 * 说明:
 *
 *
 * k 的值为正数，且总是小于给定排序数组的长度。
 * 数组不为空，且长度不超过 10^4
 * 数组里的每个元素与 x 的绝对值不超过 10^4
 *
 *
 *
 *
 * 更新(2017/9/19):
 * 这个参数 arr 已经被改变为一个整数数组（而不是整数列表）。 请重新加载代码定义以获取最新更改。
 *
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	result := list.New()

	if k == 0 {
		return []int{}
	}

	targetIndex := binarySearch(arr, x)

	left, right := targetIndex-1, targetIndex
	for k > 0 {

		if left >= 0 && right <= len(arr)-1 {
			if x-arr[left] > arr[right]-x {
				result.PushBack(arr[right])
				right++
			} else if x-arr[left] <= arr[right]-x {
				result.PushFront(arr[left])
				left--
			}
		} else {
			if left < 0 {
				result.PushBack(arr[right])
				right++
			} else if right > len(arr)-1 {
				result.PushFront(arr[left])
				left--
			}
		}
		k--
	}
	length := result.Len()
	intArray := make([]int, length)
	for i := 0; i < length; i++ {
		element := result.Front()
		intArray[i] = element.Value.(int)
		result.Remove(element)
	}

	return intArray
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	var middle int

	for left < right {
		middle = left + (right-left)/2

		if target <= arr[middle] {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

// @lc code=end
