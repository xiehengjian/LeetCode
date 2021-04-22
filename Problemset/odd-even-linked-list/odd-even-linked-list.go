
// @Title: 奇偶链表 (Odd Even Linked List)
// @Author: 846188037@qq.com
// @Date: 2021-02-01 11:04:08
// @Runtime: 4 ms
// @Memory: 3.2 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head==nil{
        return head
    }
    oddhead:=head 
    evenhead:=head.Next
    evenhead2:=evenhead
    for oddhead.Next!=nil && evenhead.Next!=nil{
        oddhead.Next=oddhead.Next.Next
        oddhead=oddhead.Next
        evenhead.Next=evenhead.Next.Next
        evenhead=evenhead.Next
    }
    oddhead.Next=evenhead2
    return head
}
