package TwoPointers

/*  复杂度
* 时间 O(n) 空间(1)
* 思路：双指针->相遇问题
* 从数组两端走起，每次迭代时判断左pointer和右pointer指向的数字哪个大，如果左pointer小，
* 意味着向左移动右pointer不可能使结果变得更好，因为瓶颈在左pointer，移动右pointer只会变小，所以这时候我们选择左pointer右移。
* 反之，则选择右pointer左移。在这个过程中一直维护最大的那个容积
*/
func maxArea(h []int) int {
	// two pointers
	left := 0
	right := len(h) - 1
	ret := 0

	for left < right {
		ret = Max(ret, Min(h[left], h[right]) * (right - left))
		if h[left] <= h[right] {
			left++
		} else {
			right--
		}
	}
	return ret
}

func Min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}