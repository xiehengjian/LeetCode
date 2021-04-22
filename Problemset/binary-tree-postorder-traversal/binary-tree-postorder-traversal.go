
// @Title: 二叉树的后序遍历 (Binary Tree Postorder Traversal)
// @Author: 846188037@qq.com
// @Date: 2021-02-05 17:42:14
// @Runtime: 0 ms
// @Memory: 2.2 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res:=make([]int,0)
    if root==nil{
        return res
    }
    stack1:=make([]*TreeNode,0)
    stack1=append(stack1,root)
    stack2:=make([]*TreeNode,0)

    for len(stack1)!=0{
        head:=stack1[len(stack1)-1]
        stack1=stack1[:len(stack1)-1]
        stack2=append(stack2,head)
        fmt.Println(head)
        if head.Left!=nil{
            stack1=append(stack1,head.Left)
        }
        if head.Right!=nil{
            stack1=append(stack1,head.Right)
        }
    }

    for len(stack2)!=0{
        res=append(res,stack2[len(stack2)-1].Val)
        stack2=stack2[:len(stack2)-1]
    }
    return res

}
