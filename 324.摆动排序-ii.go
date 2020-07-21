package main

/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 */

// @lc code=start
func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	length := len(nums)
	middle := findMiddle(&nums, 0, length-1, (length-1)/2)
	threeWayPartition(&nums, middle)
	left := make([]int, middle+1)
	right := make([]int, length-middle-1)
	copy(left, nums[:middle+1])
	copy(right, nums[middle+1:])
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			nums[i] = left[len(left)-i/2-1]
		} else {
			nums[i] = right[len(right)-i/2-1]
		}
	}
}

// [l, q) is < middle
// [k+1, r] > middle
func findMiddle(nums *[]int, l, r, rank int) int {
	pivotal := (*nums)[l]

	k := l
	current := l + 1
	for current <= r {
		if (*nums)[current] < pivotal {
			swap(nums, current, k+1)
			k++
			current++
		} else {
			current++
		}
	}
	swap(nums, l, k)

	if k == rank {
		return k
	} else if k > rank {
		return findMiddle(nums, l, k-1, rank)
	} else {
		return findMiddle(nums, k+1, r, rank)
	}
}

// [0....j) < middle
// [j,...k] == middle
// (k.....len-1] > middle
func threeWayPartition(nums *[]int, middle int) {
	middleValue := (*nums)[middle]
	swap(nums, middle, 0)
	j := 0
	k := 0

	current := 1

	for current < len(*nums) {
		if (*nums)[current] > middleValue {
			current++
		} else if (*nums)[current] == middleValue {
			swap(nums, k+1, current)
			k++
			current++
		} else {
			swap(nums, j, current)
			j++
			swap(nums, current, k+1)
			k++
			current++
		}
	}

}

func swap(nums *[]int, i, j int) {
	temp := (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = temp
}

// @lc code=end
func main() {
	nums := []int{1, 2, 2, 5, 2, 1, 5, 2, 1, 3, 4, 4, 1, 5, 2}
	wiggleSort(nums)

}
