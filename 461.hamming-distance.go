package main

import "fmt"

/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	m := x ^ y
	count := 0
	for m > 0 {
		if m&1 > 0 {
			count++
		}
		m >>= 1
	}
	return count
}

// @lc code=end
func main() {
	m := hammingDistance(2, 1)
	fmt.Printf("%d\n", m)
}
