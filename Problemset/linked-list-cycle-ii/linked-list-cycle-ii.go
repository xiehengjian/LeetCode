
// @Title: 环形链表 II (Linked List Cycle II)
// @Author: 846188037@qq.com
// @Date: 2021-02-03 14:22:50
// @Runtime: 8 ms
// @Memory: 3.7 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow:=head 
    fast:=head 
    for fast!=nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
        if fast==slow{
            break
        }
    }
    if fast==nil || fast.Next==nil{
        return nil
    }
    second:=head 
    for slow!=second{
        second=second.Next
        slow=slow.Next
    }
    return slow

    
}
