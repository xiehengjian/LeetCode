
// @Title: 从前序与中序遍历序列构造二叉树 (Construct Binary Tree from Preorder and Inorder Traversal)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 00:11:38
// @Runtime: 8 ms
// @Memory: 4.2 MB

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0{
		return nil
	}
	root:=&TreeNode{Val:preorder[0]}
	index:=0
	for i:=0;i<len(inorder);i++{
		if inorder[i]==preorder[0]{
			index=i
		}
	}
	root.Left=buildTree(preorder[1:len(inorder[:index])+1],inorder[:index])
	root.Right=buildTree(preorder[len(inorder[:index])+1:],inorder[index+1:])
	return root


}


// @lc code=end


