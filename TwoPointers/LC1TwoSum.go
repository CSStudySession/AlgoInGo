package TwoPointers

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int, len(nums))

	for i, num := range nums {
		if idx, ok := dict[target-num]; ok {
			return []int{idx, i}
		}
		dict[num] = i
	}
	return []int{0, 0}
}