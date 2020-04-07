/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	ln := len(needle)
	lh := len(haystack)
	for i := 0; i <= (lh - ln); i++ {
		if haystack[i:(i+ln)] == needle {
			return i
		}
	}
	return -1
}

// @lc code=end

