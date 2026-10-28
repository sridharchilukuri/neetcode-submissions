/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head1, head2 := list1, list2

    dummy := &ListNode{}
    curr := dummy

    for head1 != nil && head2 != nil {

        if head1.Val <= head2.Val {
            curr.Next = head1
            head1 = head1.Next
        } else {
            curr.Next = head2
            head2 = head2.Next
        }

        curr = curr.Next
    }

    if head1 != nil {
        curr.Next = head1
    } else {
        curr.Next = head2
    }

    return dummy.Next
}