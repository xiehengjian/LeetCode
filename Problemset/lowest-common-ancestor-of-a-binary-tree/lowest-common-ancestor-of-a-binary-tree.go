
// @Title: 二叉树的最近公共祖先 (Lowest Common Ancestor of a Binary Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 09:37:34
// @Runtime: 16 ms
// @Memory: 7.6 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     if root==nil{
         return nil
     }
     if root==p || root==q{
		 return root 
	 }
     left:=lowestCommonAncestor(root.Left,p,q)
     right:=lowestCommonAncestor(root.Right,p,q)
     if left!=nil && right!=nil{
         return root
     }
     if left!=nil{
         return left
     }
    if right!=nil{
        return right
    }
    return nil


  
}
