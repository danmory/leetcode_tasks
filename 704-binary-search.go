package main

import "fmt"

/*
Given an array of integers nums which is sorted in ascending order,
and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
You must write an algorithm with O(log n) runtime complexity.
*/
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for right >= left {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9) == 4)
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2) == -1)
}
