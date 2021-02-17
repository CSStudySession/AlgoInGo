package main

import "math"

func myAtoi(str string) int {
	var (
		ret int64
		start bool //  遇到第一个数字后置为true
		sign = 1 // 初始成正数
	)

	for _, char := range str { // 遍历每个字符
		if char >= '0' && char <= '9' {
			start = true
			ret = ret * 10 + int64(char - '0') // type coercion
			if ret >= math.MaxInt32 + 1 { // overflow case
				break
			}
		} else if !start && char == ' ' {
			continue
		} else if !start && char == '+' {
			start = true
		} else if !start && char == '-' {
			sign = -1
			start = true
		} else {
			break
		}
	}

	ret *= int64(sign) // 加上符号
	if ret > math.MaxInt32 {
		ret = math.MaxInt32
	}
	if ret < math.MinInt32 {
		ret = math.MinInt32
	}

	return int(ret)
}