
// @Title: 平衡二叉树 (Balanced Binary Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-14 14:38:01
// @Runtime: 16 ms
// @Memory: 5.7 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root==nil{
		return true
	}
	check:=depth(root.Left)-depth(root.Right)<=1 && depth(root.Left)-depth(root.Right)>=-1
	return check&&isBalanced(root.Left)&&isBalanced(root.Right)
}

func depth(root *TreeNode)int{
	if root==nil{
		return 0
	}
	if root.Left==nil&&root.Right==nil{
		return 1
	}
	return 1+max(depth(root.Left),depth(root.Right))
}

func max(x,y int)int{
	if x>y{
		return x
	}
	return y
}


