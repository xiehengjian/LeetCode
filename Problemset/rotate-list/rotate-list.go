
// @Title: 旋转链表 (Rotate List)
// @Author: 846188037@qq.com
// @Date: 2021-03-23 09:00:26
// @Runtime: 4 ms
// @Memory: 2.6 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head==nil || head.Next==nil{
        return head
    }
    p:=head 
    len:=1
    for p.Next!=nil{
        p=p.Next
        len++
    }
    p.Next=head 
    p=p.Next

    for i:=0;i<(len-k%len-1);i++{
        p=p.Next
    }
    temp:=p.Next
    p.Next=nil 
    return temp

}
