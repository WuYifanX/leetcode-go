package main

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 *
 * https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (42.90%)
 * Likes:    631
 * Dislikes: 0
 * Total Accepted:    66.5K
 * Total Submissions: 155K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个非空二叉树，返回其最大路径和。
 *
 * 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   3
 *
 * 输出: 6
 *
 *
 * 示例 2:
 *
 * 输入: [-10,9,20,null,null,15,7]
 *
 * -10
 * / \
 * 9  20
 * /  \
 * 15   7
 *
 * 输出: 42
 *
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	return maxOf(innerMaxPathSum(root))
}

// return (第一个最大的结果不能被修改, 这个最大的结果可以被修改)
func innerMaxPathSum(node *TreeNode) (int, int) {
	if node == nil {
		return -1000, -1000
	}

	if node.Left == nil && node.Right == nil {
		return node.Val, node.Val
	}

	leftNodeFullPathMaxSum, leftNodePartialPathMaxSum := innerMaxPathSum(node.Left)
	rightNodeFullPathMaxSum, rightNodePartialPathMaxSum := innerMaxPathSum(node.Right)

	fullPathMaxSum := maxOf(maxOf(leftNodeFullPathMaxSum, rightNodeFullPathMaxSum), node.Val+leftNodePartialPathMaxSum+rightNodePartialPathMaxSum)

	partialPathMaxSum := node.Val + maxOf(0, maxOf(leftNodePartialPathMaxSum, rightNodePartialPathMaxSum))
	fullPathMaxSum = maxOf(fullPathMaxSum, maxOf(leftNodePartialPathMaxSum, rightNodePartialPathMaxSum))
	return fullPathMaxSum, partialPathMaxSum
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
