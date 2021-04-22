
// @Title: 合并K个升序链表 (Merge k Sorted Lists)
// @Author: 846188037@qq.com
// @Date: 2021-04-20 20:21:43
// @Runtime: 104 ms
// @Memory: 5.9 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists)==0{
        return nil 
    }
    for len(lists)>1{
        lists=append(lists,merge(lists[0],lists[1]))
        lists=lists[2:]
    }
    return lists[0]
}

func merge(node1 ,node2 *ListNode)*ListNode{
    if node1==nil{
        return node2
    }
    if node2==nil{
        return node1
    }
    dummy:=&ListNode{}
    head:=dummy
    for node1!=nil && node2!=nil{
        if node1.Val<node2.Val{
            head.Next=node1
            node1=node1.Next
        }else{
            head.Next=node2
            node2=node2.Next
        }
        head=head.Next
    }
    Print(node1)
    Print(node2)
    Print(head)
    for node1!=nil{
        head.Next=node1
        node1=node1.Next
        head=head.Next//这里一开始忘了写了，很扯
    }
    for node2!=nil{
        head.Next=node2
        node2=node2.Next
        head=head.Next
    }
    Print(dummy.Next)
    return dummy.Next
}

func Print(node *ListNode){
    for node!=nil{
        fmt.Print(node.Val)
        node=node.Next
    }
    fmt.Println(" ")
}
