import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (38.34%)
 * Likes:    3950
 * Dislikes: 279
 * Total Accepted:    575.7K
 * Total Submissions: 1.5M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given a collection of intervals, merge all overlapping intervals.
 *
 * Example 1:
 *
 *
 * Input: [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into
 * [1,6].
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 *
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergeList := make([][]int, 0)
	for _, elem := range intervals {
		if len(mergeList) == 0 || mergeList[len(mergeList)-1][1] < elem[0] {
			mergeList = append(mergeList, elem)
		} else {
			maxTail := max(mergeList[len(mergeList)-1][1], elem[1])
			mergeList[len(mergeList)-1][1] = maxTail
		}
	}
	return mergeList
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

// @lc code=end

