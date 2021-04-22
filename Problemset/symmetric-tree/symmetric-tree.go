
// @Title: 对称二叉树 (Symmetric Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 10:12:58
// @Runtime: 0 ms
// @Memory: 2.9 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root==nil{
		return true
	}
	root.Right=reverse(root.Right)
	return check(root.Left,root.Right)
	
}

func reverse(root *TreeNode) *TreeNode{
	if root==nil{
		return nil
	}
	temp:=root.Left
	root.Left=root.Right
	root.Right=temp 
	root.Left=reverse(root.Left)
	root.Right=reverse(root.Right)
	return root
}

func check(root1 *TreeNode,root2 *TreeNode)bool{
	if root1==nil && root2==nil{
		return true
	}
	if root1==nil || root2==nil{
		return false
	}
	if root1.Val!=root2.Val{
		return false
	}
	if !check(root1.Left,root2.Left)||!check(root1.Right,root2.Right){
		return false
	}
	return true

}
