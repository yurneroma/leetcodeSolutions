/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

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

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				if obstacleGrid[i][j] == 0 {
					dp[i][j] = 1
				}
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				}
			} else {
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i][j-1] + dp[i-1][j]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

