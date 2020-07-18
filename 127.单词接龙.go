package main

import (
	"container/list"
)

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	visited := make(map[string]int)
	visited[beginWord] = 1
	wordMap := make(map[string]bool)

	for _, word := range wordList {
		wordMap[word] = true
	}

	queue := list.New()
	queue.PushBack(beginWord)

	for queue.Len() != 0 {
		currentWord := queue.Front().Value.(string)
		queue.Remove(queue.Front())

		for _, value := range getNexts(currentWord, &wordMap) {
			if _, exist := visited[value]; !exist {
				visited[value] = visited[currentWord] + 1
				queue.PushBack(value)
			}
		}
	}

	steps, ok := visited[endWord]
	if ok {
		return steps
	}

	return 0
}

func getNexts(from string, wordMap *map[string]bool) []string {
	nexts := make([]string, 0)
	dicts := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(from); i++ {
		for index := 0; index < 26; index++ {
			next := from[0:i] + string(dicts[index]) + from[i+1:]
			if next == from {
				continue
			}

			if _, ok := (*wordMap)[next]; ok {
				nexts = append(nexts, next)
			}
		}
	}
	return nexts
}

// @lc code=end
