package main
// link: https://leetcode.com/problems/sum-of-subarray-minimums/

func sumSubarrayMins(A []int) int {
	mod := int(1e9 + 7)
	ret := 0
	stack := make([]int, 0)

	for i := 0; i < len(A); i++ {
		for len(stack) > 0 && A[stack[len(stack) - 1]] > A[i] {
			mid := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack) - 1]
			}

			ret = (ret + A[mid] * (i - mid) * (mid - left)) % mod
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		mid := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		left := -1
		if len(stack) > 0 {
			left = stack[len(stack) - 1]
		}
		ret = (ret + A[mid] * (len(A) - mid) * (mid - left)) % mod
	}

	return ret
}
