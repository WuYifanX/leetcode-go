// package main

/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 *
 * https://leetcode-cn.com/problems/evaluate-division/description/
 *
 * algorithms
 * Medium (54.23%)
 * Likes:    183
 * Dislikes: 0
 * Total Accepted:    9.7K
 * Total Submissions: 17.9K
 * Testcase Example:  '[["a","b"],["b","c"]]\n' +
  '[2.0,3.0]\n' +
  '[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]'
 *
 * 给出方程式 A / B = k, 其中 A 和 B 均为用字符串表示的变量， k
 * 是一个浮点型数字。根据已知方程式求解问题，并返回计算结果。如果结果不存在，则返回 -1.0。
 *
 * 示例 :
 * 给定 a / b = 2.0, b / c = 3.0
 * 问题:  a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
 * 返回 [6.0, 0.5, -1.0, 1.0, -1.0 ]
 *
 * 输入为:  vector<pair<string, string>> equations, vector<double>& values,
 * vector<pair<string, string>> queries(方程式，方程式结果，问题方程式)， 其中 equations.size()
 * == values.size()，即方程式的长度与方程式结果长度相等（程式与结果一一对应），并且结果值均为正数。以上为方程式的描述。
 * 返回vector<double>类型。
 *
 * 基于上述例子，输入如下：
 *
 * equations(方程式) = [ ["a", "b"], ["b", "c"] ],
 * values(方程式结果) = [2.0, 3.0],
 * queries(问题方程式) = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x",
 * "x"] ].
 *
 *
 * 输入总是有效的。你可以假设除法运算中不会出现除数为0的情况，且不存在任何矛盾的结果。
 *
*/

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	if len(values) == 0 {
		return []float64{}
	}

	graph := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		equation := equations[i]

		_, ok := graph[equation[0]]
		if !ok {
			graph[equation[0]] = make(map[string]float64)
		}
		graph[equation[0]][equation[1]] = values[i]

		_, ok = graph[equation[1]]
		if !ok {
			graph[equation[1]] = make(map[string]float64)
		}
		graph[equation[1]][equation[0]] = 1 / values[i]
	}

	result := make([]float64, len(queries))

	for i := 0; i < len(queries); i++ {
		query := queries[i]

		_, isOk := graph[query[0]]
		if isOk && query[0] == query[1] {
			result[i] = 1.0
			continue
		}
		isVisit := make(map[string]bool)
		memo := make(map[string]float64)

		result[i] = findInGraph(&graph, query[0], query[1], &memo, &isVisit)
	}
	return result
}

func findInGraph(graph *map[string]map[string]float64, start, target string, memo *map[string]float64, isVisit *map[string]bool) float64 {
	connectedGraphs, isOk := (*graph)[start]
	_, isOk2 := (*graph)[target]
	if !isOk || !isOk2 {
		(*memo)[start+"/"+target] = -1.0
		return -1.0
	}

	for point, value := range connectedGraphs {
		if point == target {
			(*memo)[start+"/"+target] = value
			return value
		}

		if point == start {
			continue
		}

		_, isExist := (*memo)[start+"/"+point]
		if isExist {
			continue
		}

		_, isExist = (*isVisit)[start+"/"+point]
		if isExist {
			continue
		}
		(*isVisit)[start+"/"+point] = true

		result := findInGraph(graph, point, target, memo, isVisit)
		if result > 0 {
			(*memo)[start+"/"+target] = value * result
			return value * result
		}
	}
	(*memo)[start+"/"+target] = -1.0
	(*isVisit)[start] = true

	return -1.0
}

// @lc code=end

// func main() {
// 	calcEquation([][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}, []float64{3.0, 4.0, 5.0, 6.0}, [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}})
// 	calcEquation([][]string{{"a", "b"}, {"c", "d"}}, []float64{1.0, 1.0}, [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}})
// 	fmt.Println(calcEquation([][]string{{"x1", "x2"}, {"x2", "x3"}, {"x1", "x4"}, {"x2", "x5"}}, []float64{3.0, 0.5, 3.4, 5.6}, [][]string{{"x2", "x4"}, {"x1", "x5"}, {"x1", "x3"}, {"x5", "x5"}, {"x5", "x1"}, {"x3", "x4"}, {"x4", "x3"}, {"x6", "x6"}, {"x0", "x0"}}))

// }
