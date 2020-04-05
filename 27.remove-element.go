/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	i := 0
	for _, elem := range nums {
		if elem != val {
			nums[i] = elem
			i++
		}
	}
	return i
}

// @lc code=end

