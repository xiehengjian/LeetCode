
// @Title: 二叉搜索树中的插入操作 (Insert into a Binary Search Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 15:24:07
// @Runtime: 28 ms
// @Memory: 7.3 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root==nil{
		return &TreeNode{Val:val}
	}
	if root.Val<val{
		root.Right=insertIntoBST(root.Right,val)
	}
	if root.Val>val{
		root.Left=insertIntoBST(root.Left,val)
	}
	return root

}
