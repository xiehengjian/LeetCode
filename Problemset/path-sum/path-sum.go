
// @Title: 路径总和 (Path Sum)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 10:51:33
// @Runtime: 8 ms
// @Memory: 4.6 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root==nil{
		return false
	}
	if root.Val==targetSum && root.Left==nil &&root.Right==nil{
		return true
	}
	
	return hasPathSum(root.Left,targetSum-root.Val)||hasPathSum(root.Right,targetSum-root.Val)

}
