
// @Title: 反转链表 (反转链表 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 13:16:18
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
    var pre *ListNode
    for head!=nil{
        temp:=head.Next
        head.Next=pre 
        pre=head
        head=temp
    }
    return pre
}
