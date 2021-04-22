
// @Title: 两两交换链表中的节点 (Swap Nodes in Pairs)
// @Author: 846188037@qq.com
// @Date: 2021-02-02 22:47:00
// @Runtime: 0 ms
// @Memory: 2.1 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummy:=&ListNode{Val:0}
    dummy.Next=head
    head=dummy
    for head!=nil && head.Next!=nil && head.Next.Next!=nil{
        node1:=head.Next
        node2:=head.Next.Next
        head.Next=head.Next.Next
        node1.Next=node1.Next.Next
        node2.Next=node1
        head=head.Next.Next
       
    }
    return dummy.Next

}
