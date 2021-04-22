
// @Title: 相同的树 (Same Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 10:29:29
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p==nil&&q==nil{
		return true
	}
	if p==nil || q==nil{
		return false
	}
	if p.Val!=q.Val{
		return false
	}
	return isSameTree(p.Left,q.Left)&&isSameTree(p.Right,q.Right)

}
