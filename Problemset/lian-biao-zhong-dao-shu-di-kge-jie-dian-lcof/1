
// @Title: 链表中倒数第k个节点 (链表中倒数第k个节点 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 13:13:56
// @Runtime: 4 ms
// @Memory: 2.3 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
    first:=head
    second:=head 
    for i:=0;i<k;i++{
        if second!=nil{
            second=second.Next
        }else{
            return nil
        }
        
    }
    for second!=nil{
        first=first.Next
        second=second.Next
        fmt.Println(first)
    }
    return first

}
