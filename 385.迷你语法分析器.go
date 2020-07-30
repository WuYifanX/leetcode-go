package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=385 lang=golang
 *
 * [385] 迷你语法分析器
 */
type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool {}

func (n *NestedInteger) SetInteger(value int)   {}
func (n *NestedInteger) Add(elem NestedInteger) {}

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	result := &NestedInteger{}
	if len(s) == 0 {
		return result
	}
	if s[0] != '[' {
		value, _ := strconv.Atoi(s)
		result.SetInteger(value)
		return result
	}

	if len(s) <= 2 {
		return result
	}

	start := 1
	depth := 0

	for i := 1; i < len(s); i++ {
		if depth == 0 && (s[i] == ',' || i == len(s)-1) {
			result.Add(*deserialize(s[start:i]))
			start = i + 1
		} else if s[i] == '[' {
			depth++
		} else if s[i] == ']' {
			depth--
		}
	}
	return result
}

// @lc code=end
