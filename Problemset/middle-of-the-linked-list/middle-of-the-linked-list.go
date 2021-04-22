
// @Title: 链表的中间结点 (Middle of the Linked List)
// @Author: 846188037@qq.com
// @Date: 2021-03-18 10:43:47
// @Runtime: 0 ms
// @Memory: 1.9 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow:=head 
    fast:=head
    for fast!=nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
    }
    return slow


}
