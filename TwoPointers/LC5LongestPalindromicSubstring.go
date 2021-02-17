package TwoPointers

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end, longest := 0, 0, 0
	for i := 0; i < len(s); i++ {
		currLen, left, right := getLongestPali(s, i, i) // 奇数回文
		if currLen > longest {
			longest = currLen
			start = left
			end = right
		}

		if i + 1 < len(s) { // 偶数回文
			currLen, left, right = getLongestPali(s, i, i + 1)
			if currLen > longest {
				longest = currLen
				start = left
				end = right
			}
		}
	}

	return s[start: end + 1] // 左闭右开
}

func getLongestPali(s string, left int, right int) (curr, start, end int) {
	maxLen := 1
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}

	if right - left - 1 > maxLen {
		maxLen = right - left - 1
	}

	curr, start, end = maxLen, left + 1, right - 1 // [left+1, right-1]为潜在合法回文
	return
}