
// @Title: 将有序数组转换为二叉搜索树 (Convert Sorted Array to Binary Search Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-14 21:45:44
// @Runtime: 4 ms
// @Memory: 3.4 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==0{
		return nil
	}
	mid:=len(nums)/2
	root:=&TreeNode{Val:nums[mid]}
	root.Left=sortedArrayToBST(nums[:mid])
	root.Right=sortedArrayToBST(nums[mid+1:])
	return root

}


