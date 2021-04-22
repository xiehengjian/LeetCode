
// @Title: 删除链表的倒数第 N 个结点 (Remove Nth Node From End of List)
// @Author: 846188037@qq.com
// @Date: 2021-02-03 13:47:40
// @Runtime: 0 ms
// @Memory: 2.2 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy:=&ListNode{}
    dummy.Next=head
    head=dummy
    first:=head
    second:=head 
    for i:=0;i<n+1;i++{
        second=second.Next
    }
    for second!=nil{
        first=first.Next
        second=second.Next
    }
    first.Next=first.Next.Next
    return head.Next
}
