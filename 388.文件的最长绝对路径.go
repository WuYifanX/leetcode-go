package main

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=388 lang=golang
 *
 * [388] 文件的最长绝对路径
 */

//@lc code=start
func lengthLongestPath(input string) int {
	if len(input) <= 1 || !strings.Contains(input, ".") {
		return 0
	}

	inputArray := strings.Split(input, "\n")

	depth := 0
	result := 0
	lengthForCurrentPath := 0
	prevDepth := 0

	parentDepthMap := make([]int, len(inputArray)+1)
	for i := 0; i < len(inputArray); i++ {
		depth = calculateDepth(inputArray[i])
		if depth > prevDepth {
			lengthForCurrentPath += stripWords(inputArray[i])
		} else if depth < prevDepth {
			clearLength := clearArrayFromIndex(depth, parentDepthMap)
			lengthForCurrentPath = lengthForCurrentPath - clearLength + stripWords(inputArray[i])
		} else if depth == prevDepth {
			lengthForCurrentPath = lengthForCurrentPath - parentDepthMap[depth] + stripWords(inputArray[i])
		}

		if strings.Contains(inputArray[i], ".") {
			result = max(result, lengthForCurrentPath+depth-1)
		}
		parentDepthMap[depth] = stripWords(inputArray[i])
		prevDepth = depth

		// fmt.Println(inputArray[i], depth, lengthForCurrentPath, parentDepthMap, result)

	}
	return result
}

func clearArrayFromIndex(index int, array []int) int {
	clearLength := 0
	for i := index; i < len(array); i++ {
		if array[i] == 0 {
			return clearLength
		}
		clearLength += array[i]
		array[i] = 0
	}
	return clearLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stripWords(input string) int {
	if len(input) == 0 {
		return 0
	}

	for startIndex := 0; startIndex < len(input); startIndex++ {
		if input[startIndex] == 9 || input[startIndex] == 10 {
			continue
		} else {
			return len(input) - startIndex
		}
	}
	return 0
}

func calculateDepth(input string) int {
	if len(input) == 0 {
		return 0
	}
	depth := 1

	for i := 0; i < len(input); {
		// 9 is the asci code for \t
		if input[i] != 9 {
			return depth
		}
		i++
		depth++
	}
	return depth
}

// func main() {
// 	// fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1"))
// 	// fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1subsubdir1subsubdir1subsubdir1subsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
// 	// 	fmt.Println(stripWords( "\n\tsubdir1"))
// }

// @lc code=end
