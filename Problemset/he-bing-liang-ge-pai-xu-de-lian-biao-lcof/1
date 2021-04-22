
// @Title: 合并两个排序的链表 (合并两个排序的链表  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 13:20:29
// @Runtime: 4 ms
// @Memory: 4.1 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    res:=&ListNode{}
    head:=res
    for l1!=nil && l2!=nil{
        if l1.Val>l2.Val{
            head.Next=l2 
            l2=l2.Next

        }else{
            head.Next=l1
            l1=l1.Next
        }
        head=head.Next
    }
    for l1!=nil{
        head.Next=l1 
        l1=l1.Next
        head=head.Next
    }
    for l2!=nil{
        head.Next=l2 
        l2=l2.Next
        head=head.Next
    }
    return res.Next

}
