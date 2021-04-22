
// @Title: 二叉树展开为链表 (Flatten Binary Tree to Linked List)
// @Author: 846188037@qq.com
// @Date: 2021-04-10 16:36:36
// @Runtime: 4 ms
// @Memory: 2.8 MB

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	if root==nil{
		return 
	}
	flatten(root.Left)
	flatten(root.Right)
	Left:=root.Left
	Right:=root.Right
	root.Left=nil
	root.Right=Left

    p:=root
	for p.Right!=nil{
		p=p.Right
	}
	p.Right=Right
}
// @lc code=end


