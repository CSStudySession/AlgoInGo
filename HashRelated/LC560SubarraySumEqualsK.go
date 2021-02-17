// link: https://leetcode.com/problems/subarray-sum-equals-k/
package main

import "fmt"

func subarraySum(nums []int, k int) int {
	pfxSum := make(map[int]int)
	pfxSum[0] = 1
	sum := 0
	res := 0

	for _, val := range nums {
		sum += val
		if pfxSum[sum - k] > 0 {
			res += pfxSum[sum - k]
		}
		pfxSum[sum]++
	}
	return res
}


func main() {
	nums := []int{1, 1, 1}
	k := 2
	res := subarraySum(nums, k)
	fmt.Print(res)
}
