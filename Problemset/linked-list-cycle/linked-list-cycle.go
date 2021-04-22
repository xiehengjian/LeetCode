
// @Title: 环形链表 (Linked List Cycle)
// @Author: 846188037@qq.com
// @Date: 2021-02-03 14:05:23
// @Runtime: 4 ms
// @Memory: 3.6 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    slow:=head
    fast:=head 
    for fast!=nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
        if slow==fast{
            return true
        }
    }
    return false
    
}
