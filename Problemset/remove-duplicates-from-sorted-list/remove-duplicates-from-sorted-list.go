
// @Title: 删除排序链表中的重复元素 (Remove Duplicates from Sorted List)
// @Author: 846188037@qq.com
// @Date: 2021-02-03 12:39:33
// @Runtime: 4 ms
// @Memory: 3.1 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    cur:=head
    for cur!=nil{
        for cur!=nil && cur.Next != nil && cur.Val==cur.Next.Val{
        cur.Next=cur.Next.Next
    }
    cur=cur.Next
    }
    
    return head

}
