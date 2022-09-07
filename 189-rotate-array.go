package main

import "fmt"

// Given an array, rotate the array to the right by k steps, where k is non-negative.
func rotate(nums []int, k int) {
	k %= len(nums) 
	res := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, res)
}

func main() {
	toRotate := []int{1,2,3,4,5,6,7}
	rotate(toRotate, 3)
	fmt.Println(toRotate)
}
