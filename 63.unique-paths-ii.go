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
	obstacleGrid := make([][]int, 3)
	dp := []int{0, 0, 0}
	obstacleGrid[0] = dp
	dp = []int{0, 1, 0}
	obstacleGrid[1] = dp
	dp = []int{0, 0, 0}
	obstacleGrid[2] = dp
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

	//check the edge condition, if 1 in row[0]  or 1 in col[0], then row[0]  =0, and col[0] = 0
	gridFlag := false
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			gridFlag = true
		}
	}

	//
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		indp := make([]int, n)
		dp[i] = indp
		for j := 0; j < n; j++ {
			if gridFlag && i == 0 {
				dp[i][j] = 0
			}
			if !gridFlag && i == 0 {
				dp[i][j] = 1
			}
			if j == 0 && obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}
			if j == 0 && obstacleGrid[i][j] == 0 {
				dp[i][j] = 1
			}
			if i > 0 && j > 0 {
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]

				}
			}
			//fmt.Println("i,j,value", i, j, dp[i][j])

		}
	}
	return dp[m-1][n-1]
}

// @lc code=end
