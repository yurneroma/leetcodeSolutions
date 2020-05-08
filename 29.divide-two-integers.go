/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	intMax := 1<<31 - 1
	intMin := -1 << 31

	if divisor == 0 {
		return intMax
	}

	sign := 1
	absDvd := abs(dividend)
	absDvr := abs(divisor)
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = 1
	} else {
		sign = -1
	}

	count := 0
	for absDvr <= absDvd {
		innerCount := 1
		temp := absDvr
		for temp<<1 <= absDvd {
			temp <<= 1
			innerCount <<= 1
		}
		absDvd = absDvd - temp
		count += innerCount
	}

	result := sign * count
	if result > intMax {
		return intMax
	}
	if result < intMin {
		return intMin
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

// @lc code=end

