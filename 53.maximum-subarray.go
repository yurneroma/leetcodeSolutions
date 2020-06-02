/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (45.90%)
 * Likes:    7389
 * Dislikes: 341
 * Total Accepted:    976K
 * Total Submissions: 2.1M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */

// @lc code=start
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i]+sum {
			sum = nums[i]
		} else {
			sum = nums[i] + sum
		}

		if max < sum {
			max = sum
		}
	}
	return max
}

// @lc code=end

