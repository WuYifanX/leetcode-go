package main

/*
 * @lc app=leetcode.cn id=365 lang=golang
 *
 * [365] 水壶问题
 */

// @lc code=start

import (
	"container/list"
)

type State struct {
	first  int
	second int
}

func canMeasureWater(x int, y int, z int) bool {
	if x+y == z || x == z || y == z || z == 0 {
		return true
	}

	if x+y < z {
		return false
	}

	isVisit := make(map[State]bool)

	queue := list.New()
	firstState := State{0, 0}
	queue.PushBack(firstState)

	for queue.Len() != 0 {
		currentStateElement := queue.Front()
		queue.Remove(currentStateElement)
		currentState := currentStateElement.Value.(State)

		if currentState.first == z || currentState.second == z || currentState.first+currentState.second == z {

			return true
		}

		nextStates := getNextStates(currentState, x, y)

		for i := 0; i < len(nextStates); i++ {
			nextState := nextStates[i]

			if !isVisit[nextState] {
				queue.PushBack(nextState)
				isVisit[nextState] = true
			}
		}

	}

	return false
}

func getNextStates(currentState State, x, y int) []State {
	nextStates := make([]State, 0)

	// 操作1: 把x清理掉
	nextStates = append(nextStates, State{0, currentState.second})

	// 操作2: 把y清理掉
	nextStates = append(nextStates, State{currentState.first, 0})

	// 操作3: 把x装满
	nextStates = append(nextStates, State{x, currentState.second})

	// 操作4: 把y装满
	nextStates = append(nextStates, State{currentState.first, y})

	// 操作5: 把x的东西装到y
	var minValue int
	minValue = min(currentState.first, y-currentState.second)
	nextStates = append(nextStates, State{currentState.first - minValue, currentState.second + minValue})

	// 操作6: 把y的东西撞到x

	minValue = min(currentState.second, x-currentState.first)
	nextStates = append(nextStates, State{currentState.first + minValue, currentState.second - minValue})

	return nextStates
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	} else {
		return v2
	}
}

// @lc code=end
