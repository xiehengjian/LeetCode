
// @Title: 二叉搜索树中第K小的元素 (Kth Smallest Element in a BST)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 14:41:21
// @Runtime: 12 ms
// @Memory: 6.4 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var rank=0
var res=0
func kthSmallest(root *TreeNode, k int) int {
	traverse(root,k)
	rank=0
	return res

}

func traverse(root *TreeNode,k int){
	if root==nil{
		return
	}
	traverse(root.Left,k)
	rank++
	if k==rank{
		//fmt.Println(res)
		res=root.Val
		return
	}
	traverse(root.Right,k)

}
