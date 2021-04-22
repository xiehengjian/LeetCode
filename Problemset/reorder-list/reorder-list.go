
// @Title: 重排链表 (Reorder List)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 13:12:23
// @Runtime: 12 ms
// @Memory: 5.3 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head==nil{
        return
    }
    //res:=head
    mid:=getMid(head)
    mid=reverse(mid)
    for mid!=nil && head!=nil{
        temp1:=head.Next
        temp2:=mid.Next
        head.Next=mid 
        mid.Next=temp1
        head=temp1
        mid=temp2
    }
    


}

func getMid(head *ListNode)*ListNode{
    slow,fast:=head,head 
    for fast!=nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
    }
    temp:=slow.Next
    slow.Next=nil
    return temp
}
func reverse(head *ListNode)*ListNode{
    var pre *ListNode
    for head!=nil{
        temp:=head.Next
        head.Next=pre
        pre=head
        head=temp
    }
    return pre
}

