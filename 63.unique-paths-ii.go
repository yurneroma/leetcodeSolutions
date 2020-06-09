package main

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 *
 * https://leetcode.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (33.99%)
 * Likes:    1572
 * Dislikes: 234
 * Total Accepted:    283.2K
 * Total Submissions: 829.4K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 *
 * Now consider if some obstacles are added to the grids. How many unique paths
 * would there be?
 *
 *
 *
 * An obstacle and empty space is marked as 1 and 0 respectively in the grid.
 *
 * Note: m and n will be at most 100.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * Output: 2
 * Explanation:
 * There is one obstacle in the middle of the 3x3 grid above.
 * There are two ways to reach the bottom-right corner:
 * 1. Right -> Right -> Down -> Down
 * 2. Down -> Down -> Right -> Right
 *
 *
 */

func main() {
	// obstacleGrid := make([][]int, 3)
	// dp := []int{0, 0, 0}
	// obstacleGrid[0] = dp
	// dp = []int{0, 1, 0}
	// obstacleGrid[1] = dp
	// dp = []int{0, 0, 0}
	// obstacleGrid[2] = dp
	obstacleGrid := [][]int{{1}, {0}}
	uniquePathsWithObstacles(obstacleGrid)
}

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	//check the edge condition, dp[0][i]  = dp[0][i-1]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		indp := make([]int, n)
		dp[i] = indp
		for j := 0; j < n; j++ {
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end
