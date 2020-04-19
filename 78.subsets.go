/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	subSets := make([][]int, 0)
	leng := len(nums)
	for i := 0; i < leng; i++ {
		elems := genSubSets(nums, i)
		subSets = append(subSets, elems)
	}
}
func genSubSets(nums []int, i int) [][]int {

}

// @lc code=end

