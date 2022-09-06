package main

import (
	"fmt"
)

// O(n)
func squares(nums []int) []int {
	res := make([]int, len(nums))
	for i, v := range nums {
		res[i] = v * v
	}
	return res
}

// O(n)
func findLastNegNumberID(nums []int) int {
	lastNegID := -1
	for i := range nums {
		if nums[i] < 0 {
			lastNegID = i
		}
	}
	return lastNegID
}

// O(n)
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// O(n)
func mergeAbsNonDescending(negs, pos, res []int) {
	for i, j := 0, 0; i < len(negs) || j < len(pos); {
		if i < len(negs) && j < len(pos) {
			if -negs[i] < pos[j] {
				res[i+j] = -negs[i]
				i += 1
			} else {
				res[i+j] = pos[j]
				j += 1
			}
		} else if i < len(negs) {
			res[i+j] = -negs[i]
			i += 1
		} else {
			res[i+j] = pos[j]
			j += 1
		}
	}
}

/*
Given an integer array nums sorted in non-decreasing order,
return an array of the squares of each number sorted in non-decreasing order.
*/
func sortedSquares(nums []int) []int {
	lastNegID := findLastNegNumberID(nums)
	if lastNegID < 0 {
		return squares(nums)
	}
	negs := make([]int, lastNegID+1)          // only negative numbers
	pos := make([]int, len(nums)-lastNegID-1) // only positive numbers
	copy(negs, nums[:lastNegID+1])
	copy(pos, nums[lastNegID+1:])
	reverse(negs)
	res := make([]int, len(nums))
	mergeAbsNonDescending(negs, pos, res)
	return squares(res)
}

func main() {
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
