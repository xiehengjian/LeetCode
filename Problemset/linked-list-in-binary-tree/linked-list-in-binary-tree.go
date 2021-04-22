
// @Title: 二叉树中的列表 (Linked List in Binary Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 00:06:54
// @Runtime: 24 ms
// @Memory: 6.9 MB

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
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root==nil{
        return false
    }
    return help(head,root)||isSubPath(head,root.Left)||isSubPath(head,root.Right)
}

func help(head *ListNode,root *TreeNode)bool{
    if head==nil {
		return true
	}
	if root==nil{
		return false
	}
	if head.Val !=root.Val{
		return false
	}
    return help(head.Next,root.Left) || help(head.Next,root.Right)

}
