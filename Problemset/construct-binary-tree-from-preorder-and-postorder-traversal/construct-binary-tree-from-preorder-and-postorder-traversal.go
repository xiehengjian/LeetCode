
// @Title: 根据前序和后序遍历构造二叉树 (Construct Binary Tree from Preorder and Postorder Traversal)
// @Author: 846188037@qq.com
// @Date: 2021-04-12 10:27:05
// @Runtime: 8 ms
// @Memory: 3.2 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre)==0{
		return nil
	}
	rootVal:=pre[0]
	root:=&TreeNode{Val:rootVal}
	if len(pre)==1{
		return root
	}
	split:=0
	for i:=0;i<len(post);i++{
		if post[i]==pre[1]{
			split=i
		}
	}

	root.Left=constructFromPrePost(pre[1:split+2],post[:split+1])
	root.Right=constructFromPrePost(pre[split+2:],post[split+1:len(post)-1])
	return root

}

