
// @Title: 二叉树的最大深度 (Maximum Depth of Binary Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 10:36:29
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
func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	LeftDepth:=maxDepth(root.Left)
	RightDepth:=maxDepth(root.Right)
	if LeftDepth>RightDepth{
		return 1+LeftDepth
	}
	return 1+RightDepth

}
