
// @Title: 反转链表 (Reverse Linked List)
// @Author: 846188037@qq.com
// @Date: 2021-04-02 08:27:16
// @Runtime: 0 ms
// @Memory: 2.5 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head==nil{
        return head
    }
    var pre *ListNode
    for head!=nil{
        temp:=head.Next
        head.Next=pre
        pre=head
        head=temp
    }
    return pre

}
