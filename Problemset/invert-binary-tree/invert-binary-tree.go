
// @Title: 翻转二叉树 (Invert Binary Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-10 16:23:55
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
func invertTree(root *TreeNode) *TreeNode {
	if root==nil{
		return root
	}

	var temp *TreeNode
	temp=root.Left
	root.Left=root.Right
	root.Right=temp

	invertTree(root.Left)
	invertTree(root.Right)

	return root



}
