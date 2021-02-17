package LinkedLists

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry // 每次加上次的进位

		carry = sum / 10 // 给下次的进位
		sum %= 10 // 十进制下最后一位作为本次node的值

		curr.Next = &ListNode{Val : sum}
		curr = curr.Next
	}

	if carry == 1 {
		curr.Next = &ListNode{Val : carry}
	}

	return dummy.Next
}
