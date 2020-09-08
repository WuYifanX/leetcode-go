// package main

import "strings"

/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 */

// @lc code=start
func isValidSerialization(preorder string) bool {
	if len(preorder) == 0 {
		return false
	}

	if len(preorder) == 1 && preorder == "#" {
		return true
	}

	preorders := strings.Split(preorder, ",")

	if preorder[0] == '#' {
		return false
	}

	isLeftValid, leftEndIndex := isValidSerializationInner(&preorders, 1)

	if isLeftValid {
		isRightValid, rightEndIndex := isValidSerializationInner(&preorders, leftEndIndex+1)

		if isRightValid && rightEndIndex == len(preorders)-1 {
			return true
		}
	}

	return false
}

func isValidSerializationInner(preorders *[]string, index int) (bool, int) {
	if index >= len(*preorders) {
		return false, -1
	}

	if (*preorders)[index] == "#" {
		return true, index
	}

	isLeftValid, leftEndIndex := isValidSerializationInner(preorders, index+1)

	if isLeftValid {
		isRightValid, rightEndIndex := isValidSerializationInner(preorders, leftEndIndex+1)

		if isRightValid {
			return true, rightEndIndex
		}
	}
	return false, -1
}

// @lc code=end
