
// @Title: 二叉树的中序遍历 (Binary Tree Inorder Traversal)
// @Author: 846188037@qq.com
// @Date: 2021-02-05 16:57:27
// @Runtime: 0 ms
// @Memory: 2 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res:=make([]int,0)
    if root == nil{
        return res
    }
    stack:=make([]*TreeNode,0)
    head:=root
    for len(stack)!=0 || head!=nil{
        if head!=nil{
            stack=append(stack,head)
            head=head.Left
        }else{
            head=stack[len(stack)-1]
            stack=stack[:len(stack)-1]
            //fmt.Println(head)
            res=append(res,head.Val)
            head=head.Right
        }
    }
    return res

}
