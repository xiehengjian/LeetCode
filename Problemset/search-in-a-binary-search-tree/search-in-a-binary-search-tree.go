
// @Title: 二叉搜索树中的搜索 (Search in a Binary Search Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 15:03:12
// @Runtime: 28 ms
// @Memory: 7.2 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root==nil{
		return root 
	}
	if root.Val==val{
		return root 
	}
	if root.Val<val{
		return searchBST(root.Right,val)
	}
	return searchBST(root.Left,val)

}
