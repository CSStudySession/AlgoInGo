package TwoPointers

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {return len(s)}

	dict := make(map[byte]bool)
	ret := 1

	for left, right := 0, 0; left < len(s); left++ {
		for right < len(s) && dict[s[right]] == false { // 只要[left, right]里面没有重复字符 right一直右移
			dict[s[right]] = true // 记录右指针扫过的字符
			right++
		}
		ret = Max(ret, right - left) // 注意此时right停在了非法位置 合法长度为(right-left)
		if right == len(s) {return ret}

		if dict[s[left]] == true { // 左指针负责移除所见字符
			delete(dict, s[left])
		}
	}

	return ret
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}