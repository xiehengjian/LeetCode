
// @Title: 重建二叉树 (重建二叉树 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 08:50:25
// @Runtime: 4 ms
// @Memory: 4.2 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder)==0{
        return nil
    }
    root:=&TreeNode{Val:preorder[0]}
    index:=0
    for i:=0;i<len(inorder);i++{
        if preorder[0]==inorder[i]{
            index=i
        }
    }
    root.Left=buildTree(preorder[1:len(inorder[:index])+1],inorder[:index])
    root.Right=buildTree(preorder[len(inorder[:index])+1:],inorder[index+1:])
    return root


}


