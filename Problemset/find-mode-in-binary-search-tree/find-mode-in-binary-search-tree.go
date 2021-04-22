
// @Title: 二叉搜索树中的众数 (Find Mode in Binary Search Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-14 22:38:28
// @Runtime: 12 ms
// @Memory: 6.2 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


var count=0
var base=-1
var ans=[]int{}
var maxCount=0
func findMode(root *TreeNode) []int {
	traverse(root)
	count=0
	base=-1
	maxCount=0
	return ans
}



func traverse(root *TreeNode){
	if root==nil{
		return
	}
	traverse(root.Left)
	if root.Val==base{
		count++
	}else{
		
		base=root.Val
		count=1
	}
	if count>maxCount{
		ans=[]int{base}
		maxCount=count

	}else if count==maxCount{
		ans=append(ans,base)
	}
	traverse(root.Right)
}
