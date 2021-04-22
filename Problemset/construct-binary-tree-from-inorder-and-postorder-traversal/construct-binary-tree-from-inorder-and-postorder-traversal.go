
// @Title: 从中序与后序遍历序列构造二叉树 (Construct Binary Tree from Inorder and Postorder Traversal)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 00:33:11
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
 func buildTree(inorder []int,preorder []int ) *TreeNode {
  
	 
	if len(preorder)==0{
		return nil
	}
	root:=&TreeNode{Val:preorder[len(preorder)-1]}
	index:=0
	for i:=0;i<len(inorder);i++{
		if inorder[i]==preorder[len(preorder)-1]{
			index=i
		}
	}
	root.Left=buildTree(inorder[:index],preorder[:len(inorder[:index])])
	root.Right=buildTree(inorder[index+1:],preorder[len(inorder[:index]):len(preorder)-1])
	return root


}


// @lc code=end


