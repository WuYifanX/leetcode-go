package main

import "strconv"

/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	result := make([]string, 0)

	if len(nums) == 0 {
		return result
	}

	start := 0
	end := start
	for index, value := range nums {
		if index == 0 {
			start = value
			end = start
		} else {
			if nums[index]-nums[index-1] != +1 {
				result = append(result, transferListToString(start, end))
				start = nums[index]
				end = start
			} else {
				end = nums[index]
			}
		}
	}

	if end == -1 {
		result = append(result, transferListToString(start, start))
	} else {
		result = append(result, transferListToString(start, end))
	}

	return result
}

func transferListToString(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	}
	return strconv.Itoa(start) + "->" + strconv.Itoa(end)
}

// @lc code=end
