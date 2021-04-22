
// @Title: 有序链表转换二叉搜索树 (Convert Sorted List to Binary Search Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-14 21:41:59
// @Runtime: 4 ms
// @Memory: 5.8 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	return BuildTree(head,nil)
}

func BuildTree(left,right *ListNode)*TreeNode{
	if left==right{
		return nil
	}

	median:=getMedian(left,right)
	root:=&TreeNode{Val:median.Val}
	root.Left=BuildTree(left,median)
	root.Right=BuildTree(median.Next,right)
	return root 
}

func getMedian(left,right *ListNode)*ListNode{
	slow,fast:=left,left 
	for fast!=right && fast.Next!=right{
		slow=slow.Next 
		fast=fast.Next.Next
	}
	return slow

}
