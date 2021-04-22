
// @Title: 二叉搜索树的范围和 (Range Sum of BST)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 00:22:10
// @Runtime: 92 ms
// @Memory: 8.6 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var sum=0
func rangeSumBST(root *TreeNode, low int, high int) int {
    help(root,low,high)
    res:=sum
    sum=0
    return res

}

func help(root *TreeNode,low int,high int){
    if root==nil{
        return
    }
    help(root.Left,low,high)
    if root.Val>=low && root.Val<=high{
        sum+=root.Val
    }
    help(root.Right,low,high)
}
