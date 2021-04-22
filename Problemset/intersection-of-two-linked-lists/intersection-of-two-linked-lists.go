
// @Title: 相交链表 (Intersection of Two Linked Lists)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 18:38:47
// @Runtime: 40 ms
// @Memory: 7.2 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    node1,node2:=headA,headB
    for {
        for node1!=nil && node2!=nil{
            if node1==node2{
                return node1
            }
            node1=node1.Next
            node2=node2.Next
        }
        if node1==nil && node2==nil{
            return nil
        }
        if node1==nil{
            node1=headB
            node2=node2.Next
        }
        if node2==nil{
            node2=headA
            node1=node1.Next
        }
    }
    return nil
    
}
