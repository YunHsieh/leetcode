package main

var tmp *ListNode

func isPalindrome(head *ListNode) bool {
	tmp = head

	return check(head)
}

func check(p *ListNode) bool {
	if p == nil {
		return true
	}

	isPal := false
	if check(p.Next) && tmp.Val == p.Val {
		isPal = true
	}
	tmp = tmp.Next

	return isPal
}
