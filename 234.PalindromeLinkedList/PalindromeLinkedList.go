/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    data := []int{}
    next := head
    for {
        data = append(data, next.Val)
        if next = next.Next; next == nil {
            break
        }
    }
    mid := int(len(data)/2)
    h, t := mid-1, mid+1
    if len(data) % 2 == 0 {
        h, t = mid-1, mid
    }
    for {
        if h < 0 {
            return true
        }
        if data[h] != data[t] {
            return false
        }
        h--
        t++
    }
}
