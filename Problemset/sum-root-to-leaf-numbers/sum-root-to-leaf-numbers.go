
// @Title: 求根节点到叶节点数字之和 (Sum Root to Leaf Numbers)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 09:37:34
// @Runtime: 4 ms
// @Memory: 2.5 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    res:=[][]int{}
    temp:=[]int{}
    traverse(root,&res,temp)
    fmt.Println(res)
    sum:=0
    for _,v:=range res{
        for i:=0;i<len(v);i++{
            sum+=v[i]*int(math.Pow10(len(v)-i-1))
        }
    }
    return sum
// 1 2 3
// len=3
// i=0,10*2
//i=1,10*1
}

func traverse(root *TreeNode,res *[][]int,temp []int){
    if root==nil{
        return
    }
    temp=append(temp,root.Val)
    
    if root.Left==nil&&root.Right==nil{
        fmt.Println(temp)
        ans:=make([]int,len(temp))
        copy(ans,temp)
        *res=append(*res,ans)
        fmt.Println(*res)
        return
    }
    
    traverse(root.Left,res,temp)
   // temp=temp[:len(temp)-1]
    traverse(root.Right,res,temp)
    //temp=temp[:len(temp)-1]
}
