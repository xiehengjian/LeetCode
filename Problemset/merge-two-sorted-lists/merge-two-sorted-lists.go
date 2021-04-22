
// @Title: 合并两个有序链表 (Merge Two Sorted Lists)
// @Author: 846188037@qq.com
// @Date: 2021-02-03 13:52:56
// @Runtime: 4 ms
// @Memory: 2.5 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    p1,p2:=l1,l2 
    dummy:=&ListNode{}
    head:=dummy
    for p1!=nil && p2!=nil{
        if p1.Val<p2.Val{
            dummy.Next=p1
            p1=p1.Next
        }else{
            dummy.Next=p2
            p2=p2.Next
        }
        dummy=dummy.Next
    }
    for p1!=nil{
        dummy.Next=p1
        p1=p1.Next
        dummy=dummy.Next
    }
        for p2!=nil{
        dummy.Next=p2
        p2=p2.Next
        dummy=dummy.Next
    }
    return head.Next

}
