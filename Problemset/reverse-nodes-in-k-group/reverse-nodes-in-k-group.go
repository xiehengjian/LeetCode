
// @Title: K 个一组翻转链表 (Reverse Nodes in k-Group)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 18:14:44
// @Runtime: 28 ms
// @Memory: 4.1 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head==nil {
        return head
    }
    endPoint:=head
    cur:=head
    count:=0
    newHead:=head
    var tail *ListNode
    for endPoint!=nil{
        endPoint=endPoint.Next
        count++
        fmt.Println(count)
        if count%k==0{
            var pre *ListNode
            for cur!=endPoint{
                fmt.Println(cur.Val)
                temp:=cur.Next
                cur.Next=pre 
                pre=cur 
                cur=temp
            }
            if tail!=nil{
                tail.Next=pre
                Print(newHead)
            }else{
                newHead=pre
                tail=head
            }
            for tail!=nil &&tail.Next!=nil{
                tail=tail.Next
            }
        }
    }
    if tail!=nil{
        tail.Next=cur//反转多次则链接
    }
    
    return newHead
}


func Print(head *ListNode){
    for head!=nil{
        fmt.Print(head.Val)
        head=head.Next
    }
    fmt.Println(" ")
}
