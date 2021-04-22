
// @Title: 二叉树的镜像 (二叉树的镜像  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 13:30:35
// @Runtime: 0 ms
// @Memory: 2.1 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
    if root==nil{
        return root
    }
    temp:=root.Left
    root.Left=root.Right
    root.Right=temp
    root.Left=mirrorTree(root.Left)
    root.Right=mirrorTree(root.Right)
    return root

}
