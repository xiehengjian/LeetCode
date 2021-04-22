
// @Title: 删除链表的节点 (删除链表的节点 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 11:40:51
// @Runtime: 4 ms
// @Memory: 2.8 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
    if head==nil{
        return head
    }
    dummy:=&ListNode{}
    dummy.Next=head 
    head=dummy
    for head!=nil && head.Next!=nil{
        if head.Next.Val==val{
            head.Next=head.Next.Next
            break
        }
        head=head.Next
    }
    return dummy.Next

}
