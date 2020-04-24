/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	count := len(nums)
	hashTab := make(map[int]int, 0)
	for _, elem := range nums {
		val, _ := hashTab[elem]
		val++
		hashTab[elem] = val
	}

	for k, val := range hashTab {
		if val > count/2 {
			return k
		}
	}
	return 0
}

// @lc code=end

