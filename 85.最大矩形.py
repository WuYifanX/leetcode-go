#
# @lc app=leetcode.cn id=85 lang=python3
#
# [85] 最大矩形
#
# https://leetcode-cn.com/problems/maximal-rectangle/description/
#
# algorithms
# Hard (47.32%)
# Likes:    532
# Dislikes: 0
# Total Accepted:    36.8K
# Total Submissions: 77.7K
# Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
#
# 给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
#
# 示例:
#
# 输入:
# [
# ⁠ ["1","0","1","0","0"],
# ⁠ ["1","0","1","1","1"],
# ⁠ ["1","1","1","1","1"],
# ⁠ ["1","0","0","1","0"]
# ]
# 输出: 6
#
#

# @lc code=start
class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        if len(matrix) == 0:
            return 0

        heights = [] + [0]*(len(matrix[0])+1)
        maxArea = 0
        for row in matrix:
            stack = []
            for index, _ in enumerate(heights):
                if index < len(matrix[0]):
                    if row[index] == '1':
                        heights[index] += int(row[index])
                    else:
                        heights[index] = 0
                while len(stack) > 0 and heights[index] < heights[stack[-1]]:
                    poppedIndex = stack.pop()

                    height = heights[poppedIndex]
                    width = index - poppedIndex
                    maxArea = max(maxArea, height * width)
                    print(height, width, maxArea)

                stack.append(index)
            print(heights)

        return maxArea
# @lc code=end


# [["1","0","1","0","0"],
#  ["1","0","1","1","1"],
#  ["1","1","1","1","1"],
#  ["1","0","0","1","0"]]

3 1 3 2 2 0
1 2
