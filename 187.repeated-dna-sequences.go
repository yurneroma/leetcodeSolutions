/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 *
 * https://leetcode.com/problems/repeated-dna-sequences/description/
 *
 * algorithms
 * Medium (38.13%)
 * Likes:    643
 * Dislikes: 248
 * Total Accepted:    152.6K
 * Total Submissions: 400.3K
 * Testcase Example:  '"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"'
 *
 * All DNA is composed of a series of nucleotides abbreviated as A, C, G, and
 * T, for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to
 * identify repeated sequences within the DNA.
 *
 * Write a function to find all the 10-letter-long sequences (substrings) that
 * occur more than once in a DNA molecule.
 *
 * Example:
 *
 *
 * Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
 *
 * Output: ["AAAAACCCCC", "CCCCCAAAAA"]
 *
 *
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	return solutionWithBitAndHash(s)

}

//位操作结合哈希
//bit with hash table

func solutionWithBitAndHash(s string) []string {
	charMap := make(map[string]int, 0)
	charMap["A"] = 0
	charMap["C"] = 1
	charMap["G"] = 2
	charMap["T"] = 3
	bitStrMap := make(map[int]int, 0)
	strList := make([]string, 0)
	strMap := make(map[int]string)
	for i := 0; i < len(s)-9; i++ {
		bitStr := 0
		val := 0
		for j := i; j < i+10; j++ {
			bitStr <<= 2
			bitStr = bitStr | charMap[string(s[j])]
			bitStr = bitStr & 0xffffffff
		}
		val, ok := bitStrMap[bitStr]
		if ok {
			val++
			bitStrMap[bitStr] = val
		} else {
			bitStrMap[bitStr] = 1
		}

		if val > 1 {
			strMap[bitStr] = s[i : i+10]
		}
	}
	for _, v := range strMap {
		strList = append(strList, v)

	}

	return strList
}

// @lc code=end

