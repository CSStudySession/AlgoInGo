package main

import "math"

// Given a 32-bit signed integer, reverse digits of an integer.
func reverse(x int) int {
	ret := 0
	for x != 0 {
		ret = ret * 10 + x % 10 // 每次十进制右移一位:腾出个位给x%10
		x /= 10 // 扔掉当前最后一位:十进制下右移一位
	}

	if ret < math.MinInt32 || ret > math.MaxInt32 { // handle overflow and underflow cases
		return 0
	}

	return ret
}