
// @Title: 排序链表 (Sort List)
// @Author: 846188037@qq.com
// @Date: 2020-12-30 21:58:15
// @Runtime: 36 ms
// @Memory: 7.2 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    return mergeSort(head)

}

func findMiddle(head *ListNode)*ListNode{
    slow:=head
    fast:=head.Next
    for fast!=nil && fast.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next
    }
    return slow
}

func mergeTwoLists(l1 *ListNode,l2 *ListNode)*ListNode{
    dummy:=&ListNode{Val:0}
    head:=dummy

    for l1 !=nil && l2 !=nil{
        if l1.Val<l2.Val{
            head.Next=l1 
            l1=l1.Next
        }else{
            head.Next=l2 
            l2=l2.Next
        }
        head=head.Next
    }

    for l1!=nil{
        head.Next=l1
        head=head.Next
        l1=l1.Next
    }

    for l2!=nil{
        head.Next=l2
        head=head.Next
        l2=l2.Next
    }
    return dummy.Next
}
func mergeSort(head *ListNode)*ListNode{
    if head==nil || head.Next==nil{
        return head
    }
    middle:=findMiddle(head)
    tail:=middle.Next
    middle.Next=nil 
    left:=mergeSort(head)
    right:=mergeSort(tail)
    result:=mergeTwoLists(left,right)
    return result
}
