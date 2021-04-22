
// @Title: 删除排序链表中的重复元素 II (Remove Duplicates from Sorted List II)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 21:30:45
// @Runtime: 4 ms
// @Memory: 3 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head==nil{
        return head
    }
    dummy:=&ListNode{Val:-1}
    dummy.Next=head
    head=dummy
    base:=0
    for head!=nil && head.Next!=nil && head.Next.Next!=nil{
        if head.Next.Val==head.Next.Next.Val{
            base=head.Next.Val
            for head.Next!=nil && head.Next.Val==base{
            head.Next=head.Next.Next
        }
        }else{
        
      
            head=head.Next
        }
        
    }
    return dummy.Next

}
