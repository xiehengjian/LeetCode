
// @Title: 反转链表 II (Reverse Linked List II)
// @Author: 846188037@qq.com
// @Date: 2021-04-20 21:35:38
// @Runtime: 0 ms
// @Memory: 2.1 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if left==right || head==nil{
        return head
    }
    dummy:=&ListNode{}
    dummy.Next=head
    head=dummy

    leftTail:=head
    for i:=1;i<left;i++{
        leftTail=leftTail.Next
    }
    begin:=leftTail.Next
    rightHead:=head.Next
    for i:=0;i<right;i++{
        rightHead=rightHead.Next
    }
    cur:=leftTail.Next
    var pre *ListNode
    for cur!=rightHead{
        temp:=cur.Next
        cur.Next=pre
        pre=cur
        cur=temp
    }
    leftTail.Next=pre
    begin.Next=rightHead
    return head.Next
}
