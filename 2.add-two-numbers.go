/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ll3 := &ListNode{}
    ret := ll3
    prev := ll3
    ll1 := l1
    ll2 := l2
    rem := 0
    for ll1 != nil && ll2 != nil {
        n1 := ll1.Val
        n2 := ll2.Val
        n3 := (rem + n1 + n2) % 10
        rem = (rem + n1 + n2) / 10
        ll3.Val = n3
        ll3.Next = &ListNode{}
        ll1 = ll1.Next
        ll2 = ll2.Next
        prev = ll3
        ll3 = ll3.Next
    }
    if ll1 != nil {
        for ll1 != nil {
            n1 := ll1.Val
            n3 := (rem + n1) % 10
            rem  = (rem + n1) / 10
            ll3.Val = n3
            ll3.Next = &ListNode{}
            ll1 = ll1.Next
            prev = ll3
            ll3 = ll3.Next
        }
    } else if ll2 != nil {
        for ll2 != nil {
            n1 := ll2.Val
            n3 := (rem + n1) % 10
            rem  = (rem + n1) / 10
            ll3.Val = n3
            ll3.Next = &ListNode{}
            ll2 = ll2.Next
            prev = ll3
            ll3 = ll3.Next
        }
    }
    if rem != 0 {
        ll3.Val = rem
    } else {
        prev.Next = nil
    }
    return ret
}