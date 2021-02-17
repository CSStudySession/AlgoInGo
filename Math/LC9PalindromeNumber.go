package main

func isPalindrome(x int) bool {
	// several special cases
	if x < 0 {
		return false
	} else if x <= 9 {
		return true
	} else if x % 10 == 0 { // 最后一位是0
		return false
	}

	curr := 0
	tmp := x
	for tmp > 0 { // reverse each digit
		curr = curr * 10 + tmp % 10
		tmp /= 10
	}

	return curr == x
}