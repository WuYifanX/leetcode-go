import "container/list"

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (43.17%)
 * Likes:    361
 * Dislikes: 0
 * Total Accepted:    132.4K
 * Total Submissions: 298.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最小深度  2.
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	depth := 0
	var size int
	for queue.Len() != 0 {
		depth++
		size = queue.Len()
		for size != 0 {
			element := queue.Front()
			queue.Remove(element)
			elementTree := element.Value.(*TreeNode)
			if elementTree.Left == nil && elementTree.Right == nil {
				return depth
			}

			if elementTree.Left != nil {
				queue.PushBack(elementTree.Left)
			}

			if elementTree.Right != nil {
				queue.PushBack(elementTree.Right)
			}
			size--
		}
	}

	return -1
}

// @lc code=end
