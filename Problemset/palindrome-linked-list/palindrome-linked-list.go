
// @Title: 回文链表 (Palindrome Linked List)
// @Author: 846188037@qq.com
// @Date: 2021-03-06 10:40:30
// @Runtime: 272 ms
// @Memory: 12 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    slow:=head 
    fast:=head 
    for fast!=nil && fast.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
    }
    //mid:=reverse(slow.Next)
    mid:=reverse(slow)
    slow.Next=nil
    for head!=nil && mid!=nil{
        fmt.Print(head.Val)
        fmt.Println(mid.Val)
        if head.Val!=mid.Val{
            return false
        }
        head=head.Next
        mid=mid.Next
    }
    return true
}

func reverse(head *ListNode) *ListNode{
    var pre *ListNode
    for head != nil{
        temp:=head.Next
        head.Next=pre
        pre=head
        head=temp
    }
    return pre
}
