
// @Title: 两数相加 (Add Two Numbers)
// @Author: 846188037@qq.com
// @Date: 2021-04-22 00:54:56
// @Runtime: 16 ms
// @Memory: 4.7 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy:=&ListNode{}
    head:=dummy
    flag:=0
    for l1!=nil && l2!=nil{
       head.Next=&ListNode{Val:(l1.Val+l2.Val+flag)%10}
        if (l1.Val+l2.Val+flag)>9{
            flag=1
        }else{
            flag=0
        }
        l1=l1.Next
        l2=l2.Next
        head=head.Next
    }
    for l1!=nil{
        head.Next=&ListNode{Val:(l1.Val+flag)%10}
        if (l1.Val+flag)>9{
            flag=1
        }else{
            flag=0
        }
        l1=l1.Next
        
        head=head.Next
    }

        for l2!=nil{
        head.Next=&ListNode{Val:(l2.Val+flag)%10}
        if (l2.Val+flag)>9{
            flag=1
        }else{
            flag=0
        }
        l2=l2.Next
        
        head=head.Next
    }
    if flag==1{
       head.Next=&ListNode{Val:flag} 
    }

    return dummy.Next

}
