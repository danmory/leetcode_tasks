package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// Mock function
func isBadVersion(version int) bool {
	switch version {
	case 3:
		return false
	case 4:
		return true
	case 5:
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	right := n
	left := 1
	for {
		mid := (right + left) / 2
		if right == left {
			return mid
		}
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
}

func main() {
	fmt.Println(firstBadVersion(5) == 4)
}
