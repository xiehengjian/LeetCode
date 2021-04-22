
// @Title: 树的子结构 (树的子结构  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 13:27:10
// @Runtime: 24 ms
// @Memory: 7 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A==nil || B==nil{
        return false 
    }
    return help(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
    
}

func help(A,B *TreeNode) bool{
    if B==nil{
        return true
    }
    if A==nil{
        return false
    }
    if A.Val!=B.Val{
        return false
    }
    return help(A.Left,B.Left) && help(A.Right,B.Right)

}
