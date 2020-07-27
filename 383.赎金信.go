package main

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	dict := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		count, isExist := dict[magazine[i]]

		if isExist {
			dict[magazine[i]] = count + 1
		} else {
			dict[magazine[i]] = 1
		}
	}

	for j := 0; j < len(ransomNote); j++ {
		count, isExist := dict[ransomNote[j]]

		if isExist && count > 0 {
			dict[ransomNote[j]] = count - 1
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
