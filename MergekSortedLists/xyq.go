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

    arr := []*ListNode{l1, l2}

    res := mergeKLists2(arr);
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
func mergeKLists(lists []*ListNode) *ListNode {
    l_len := len(lists)
    if l_len < 1 {
        return nil
    }

    res, lists := lists[0], lists[1:]
    l_len--

    for l_len > 0 {
        tmp := lists[0]
        lists = lists[1:]
        res = mergeTwoLists(res, tmp)
        l_len--
    }

    return res
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists2(lists []*ListNode) *ListNode {
    if len(lists) < 1 {
        return nil
    }

    for len(lists) > 1 {
        tmp_lists := []*ListNode{}
        for len(lists) > 0 {
            // 数组里只有一个元素
            if len(lists) == 1 {
                tmp_lists = append(tmp_lists, lists[0])
                break
            }

            tmp_lists = append(tmp_lists, mergeTwoLists(lists[0], lists[1]))

            lists = lists[2:]
        }

        lists = tmp_lists

    }

    return lists[0]
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
