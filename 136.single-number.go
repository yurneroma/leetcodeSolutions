/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {

	result := 0
	for _, elem := range nums {
		result = result ^ elem
	}
	return result
}

// //solution1
// func listOperation(nums []int) int {
// 	noDupMap := make(map[int]bool, 0)
// 	keysList := make([]int, 0)
// 	for _, elem := range nums {
// 		_, ok := noDupMap[elem]
// 		if ok {
// 			delete(noDupMap, elem)
// 		} else {
// 			noDupMap[elem] = true
// 		}
// 	}
// 	for key := range noDupMap {
// 		keysList = append(keysList, key)
// 	}
// 	return keysList[0]
// }

// solution 2, with math. 2*sum(a,b,c)-sum(2a,2b,c) == c
// func withMath(nums []int) int {

// }

// //solution 3 , with xor
// func solution3(nums []int) int {
// 	result := 0
// 	for _, elem := range nums {
// 		result = result ^ elem
// 	}
// 	return result
// }

// @lc code=end

