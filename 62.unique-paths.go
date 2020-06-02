/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 *
 * https://leetcode.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (51.94%)
 * Likes:    2867
 * Dislikes: 195
 * Total Accepted:    438.5K
 * Total Submissions: 836.2K
 * Testcase Example:  '3\n2'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 *
 * How many possible unique paths are there?
 *
 *
 * Above is a 7 x 3 grid. How many possible unique paths are there?
 *
 *
 * Example 1:
 *
 *
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation:
 * From the top-left corner, there are a total of 3 ways to reach the
 * bottom-right corner:
 * 1. Right -> Right -> Down
 * 2. Right -> Down -> Right
 * 3. Down -> Right -> Right
 *
 *
 * Example 2:
 *
 *
 * Input: m = 7, n = 3
 * Output: 28
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= m, n <= 100
 * It's guaranteed that the answer will be less than or equal to 2 * 10 ^ 9.
 *
 *
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	return dp(m, n)
}

//math 排列组合
func permutation(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	result := 1
	temp := 1
	if m > n {
		m, n = n, m
	}
	//A(m-1,m+n-2)
	for i := n; i <= m+n-2; i++ {
		result *= i
	}
	for i := 1; i <= m-1; i++ {
		temp *= i
	}
	return result / temp
}

//dp
func dp(m, n int) int {
	ndp := make([][]int, m)
	for i := 0; i < m; i++ {
		ndp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				ndp[i][j] = 1
			} else {
				ndp[i][j] = ndp[i][j-1] + ndp[i-1][j]
			}
		}
	}
	return ndp[m-1][n-1]
}

// @lc code=end

