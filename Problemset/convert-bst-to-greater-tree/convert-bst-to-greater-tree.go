
// @Title: 把二叉搜索树转换为累加树 (Convert BST to Greater Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 14:41:03
// @Runtime: 12 ms
// @Memory: 6.8 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	traverse(root)
	sum=0
	return root

}
var sum=0

func traverse(root *TreeNode){
	if root==nil{
		return
	}
	traverse(root.Right)
	sum+=root.Val
	root.Val=sum

	traverse(root.Left)

}
