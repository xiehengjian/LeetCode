
// @Title: 二叉树的最小深度 (Minimum Depth of Binary Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 11:06:07
// @Runtime: 256 ms
// @Memory: 19.7 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	if root.Left==nil && root.Right==nil{
		return 1
	}
	LeftDepth:=math.MaxInt32
	RightDepth:=math.MaxInt32
	if root.Left!=nil{
		LeftDepth=min(LeftDepth,minDepth(root.Left))
	}
	if root.Right!=nil{
		RightDepth=min(RightDepth,minDepth(root.Right))
	}
	return 1+min(RightDepth,LeftDepth)

}

func min(x int,y int)int{
	if x>y{
		return y
	}
	return x
}
