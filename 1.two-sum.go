/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for index, num := range nums {
		r := target - num
		if _, ok := hash[r]; ok {
			return []int{hash[r], index}
		}
		hash[num] = index
	}
	return []int{}
}

// @lc code=end

