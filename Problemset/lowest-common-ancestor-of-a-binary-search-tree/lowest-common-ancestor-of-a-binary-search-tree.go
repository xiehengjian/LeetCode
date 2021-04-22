
// @Title: 二叉搜索树的最近公共祖先 (Lowest Common Ancestor of a Binary Search Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 08:07:53
// @Runtime: 24 ms
// @Memory: 6.9 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root==nil{
        return nil 
    }
    if (root.Val>=q.Val && root.Val<=p.Val)||(root.Val<=q.Val && root.Val>=p.Val){
        return root
    }
    if root.Val>q.Val{
        return lowestCommonAncestor(root.Left,p,q)
    }
    return lowestCommonAncestor(root.Right,p,q)
	
}
