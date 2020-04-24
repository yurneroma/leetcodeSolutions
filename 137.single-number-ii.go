/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	ans := 0
	for i := 0; i < 64; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if (nums[j] >> i & 1) == 1 {
				count++
			}
		}
		ans |= (count % 3) << i
	}
	return ans
}

// @lc code=end

