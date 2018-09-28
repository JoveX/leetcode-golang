package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if nil == head.Next {
        return nil
    }
    list := head
    del_node := &ListNode{}
    del_node = nil
    i := 0

    for nil != list {
        if i >= n && nil == del_node {
            del_node = head
        } else if i >= n {
            del_node = del_node.Next
        }
        list = list.Next
        i++
    }

    if del_node == nil {
        return head.Next
    }

    del_node.Next = del_node.Next.Next

    return head
}

func main() {
    l1 := &ListNode{Val: 1}
    l1.Next = &ListNode{Val: 2}
    l1.Next.Next = &ListNode{Val: 3}
    l1.Next.Next.Next = &ListNode{Val: 4}

    res := removeNthFromEnd(l1, 4);
    for nil != res {
        fmt.Println(res.Val)
        res = res.Next
    }
}
