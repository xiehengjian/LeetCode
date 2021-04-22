
// @Title: 从尾到头打印链表 (从尾到头打印链表 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 00:42:20
// @Runtime: 0 ms
// @Memory: 3.5 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    if head==nil{
        return []int{}
    }
    res:=[]int{}
    for head!=nil{
        res=append(res,head.Val)
        head=head.Next
    }
    ans:=[]int{}
    for i:=len(res)-1;i>-1;i--{
        ans=append(ans,res[i])
    }
    return ans

}
