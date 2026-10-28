/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    f1, s1 := head, head

    for f1 != nil && f1.Next != nil {
        f1 = f1.Next.Next
        s1 = s1.Next

        if f1 == s1 {
            return true
        }
    }

    return false
}
