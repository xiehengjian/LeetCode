
// @Title: 分隔链表 (Partition List)
// @Author: 846188037@qq.com
// @Date: 2021-02-03 15:04:58
// @Runtime: 0 ms
// @Memory: 2.4 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    bigger:=&ListNode{}
    bigHead:=bigger
    smaller:=&ListNode{}
    smallHead:=smaller

    cur:=head 
    for cur!=nil{
        if cur.Val<x{
            smallHead.Next=cur
            smallHead=smallHead.Next
        }else{
            bigHead.Next=cur
            bigHead=bigHead.Next
        }
        cur=cur.Next
    }
    bigHead.Next=nil
    smallHead.Next=bigger.Next
    return smaller.Next

}
