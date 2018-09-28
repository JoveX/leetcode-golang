package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    l1 := &ListNode{Val: 1}
    l1.Next = &ListNode{Val: 2}
    l1.Next.Next = &ListNode{Val: 4}

    l2 := &ListNode{Val: 1}
    l2.Next = &ListNode{Val: 3}
    l2.Next.Next = &ListNode{Val: 4}

    res := mergeTwoLists(l1, l2);
    for nil != res {
        fmt.Println(res.Val)
        res = res.Next
    }
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    tmp := res

    for nil != l1 || nil != l2 {
        if nil != l1 && nil != l2 {
            if l1.Val < l2.Val {
                tmp.Next = &ListNode{Val: l1.Val}
                l1 = l1.Next
            } else {
                tmp.Next = &ListNode{Val: l2.Val}
                l2 = l2.Next
            }
        } else if nil != l1 {
            tmp.Next = &ListNode{Val: l1.Val}
            l1 = l1.Next
        } else if nil != l2 {
            tmp.Next = &ListNode{Val: l2.Val}
            l2 = l2.Next
        }

        tmp = tmp.Next
    }

    return res.Next
}
