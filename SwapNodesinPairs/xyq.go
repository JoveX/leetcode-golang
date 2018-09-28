package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    l1 := &ListNode{Val: 1}
    l1.Next = &ListNode{Val: 2}
    l1.Next.Next = &ListNode{Val: 3}
    l1.Next.Next.Next = &ListNode{Val: 4}

    res := swapPairs(l1);
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
func swapPairs(head *ListNode) *ListNode {

    if head == nil || head.Next == nil {
        return head
    }

    res := head.Next
    pre := &ListNode{}

    for head != nil && head.Next != nil {
        tmp := head.Next
        head.Next = head.Next.Next
        tmp.Next = head
        if pre.Next != nil {
            pre.Next = tmp
        }
        pre = head
        head = head.Next
    }

    return res
}
