//headの後ろにListNodeをリンクしていく

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil {
        return nil
    } else if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }

    head := &ListNode{0, nil}
    curr := head

    for {
        if l1 == nil || l2 == nil {
            break
        }

        if l1.Val < l2.Val {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l2.Next
        }
        curr = curr.Next
    }

    if l1 != nil {
        curr.Next = l1
    } else {
        curr.Next = l2
    }
    return head.Next
}
