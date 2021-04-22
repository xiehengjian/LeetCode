
// @Title: 左叶子之和 (Sum of Left Leaves)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 08:42:35
// @Runtime: 0 ms
// @Memory: 2.6 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    if root==nil{
        return 0
    }
    if root.Left==nil && root.Right==nil{
        return 0
    }
    leftQueue:=[]*TreeNode{root}
    rightQueue:=[]*TreeNode{}
    sum:=0
    for len(leftQueue)!=0 || len(rightQueue)!=0{
        for len(leftQueue)!=0{
            node:=leftQueue[0]
            leftQueue=leftQueue[1:]
            if node.Left==nil && node.Right==nil{
                sum+=node.Val
            }
            if node.Left!=nil{
                leftQueue=append(leftQueue,node.Left)
            }
            if node.Right!=nil{
                rightQueue=append(rightQueue,node.Right)
            }
        }
        for len(rightQueue)!=0{
            node:=rightQueue[0]
            rightQueue=rightQueue[1:]
             if node.Left!=nil{
                leftQueue=append(leftQueue,node.Left)
            }
            if node.Right!=nil{
                rightQueue=append(rightQueue,node.Right)
            }

        }
    }
    return sum

}
