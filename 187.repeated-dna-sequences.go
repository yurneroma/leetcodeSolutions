/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	hashTab := make(map[string]int, 0)
	for i := 0; i <= len(s)-10; i++ {
		elem := s[i : i+10]
		val, _ := hashTab[elem]
		val++
		hashTab[elem] = val
	}

	strList := make([]string, 0)
	for k, v := range hashTab {
		if v > 1 {
			strList = append(strList, k)
		}
	}
	return strList
}

// @lc code=end

