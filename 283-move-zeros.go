package main

import "fmt"

func firstNonZeroID(nums []int, startingID int) int {
	toSearch := nums[startingID:]
	for i := range toSearch {
		if toSearch[i] != 0 {
			return startingID + i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
func moveZeroes(nums []int) {
	lastFoundNonZeroID := 0
	for i := range nums {
		if nums[i] == 0 {
			lastFoundNonZeroID = firstNonZeroID(nums, max(lastFoundNonZeroID, i))
			if lastFoundNonZeroID == -1 { // all next values are zeros
				return
			}
			nums[i], nums[lastFoundNonZeroID] = nums[lastFoundNonZeroID], nums[i]
		}
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)
}
