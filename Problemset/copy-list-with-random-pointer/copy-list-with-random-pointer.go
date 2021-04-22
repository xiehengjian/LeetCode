
// @Title: 复制带随机指针的链表 (Copy List with Random Pointer)
// @Author: 846188037@qq.com
// @Date: 2021-01-06 15:21:23
// @Runtime: 0 ms
// @Memory: 3.4 MB

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head==nil{
        return head 
    }

    cur:=head 
    for cur !=nil{
        clone:=&Node{Val:cur.Val,Next:cur.Next}
        temp:=cur.Next
        cur.Next=clone 
        cur=temp 
    }

    cur=head 
    for cur!=nil{
        if cur.Random!=nil{
            cur.Next.Random=cur.Random.Next
        }
        cur=cur.Next.Next
    }

    cur=head
    cloneHead:=cur.Next
    for cur!=nil&&cur.Next!=nil{
        temp:=cur.Next
        cur.Next=cur.Next.Next
        cur=temp
    }

    return cloneHead
    
}
