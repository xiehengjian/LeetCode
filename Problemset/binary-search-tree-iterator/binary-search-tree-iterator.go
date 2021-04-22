
// @Title: 二叉搜索树迭代器 (Binary Search Tree Iterator)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 11:20:40
// @Runtime: 36 ms
// @Memory: 9.8 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    Val []int

}


func Constructor(root *TreeNode) BSTIterator {
    list:=[]int{}
    helpConstructor(root,&list)
   return BSTIterator{Val:list}
}

func helpConstructor(root *TreeNode,list *[]int){
     if root==nil{
        return 
     }
    
    
    helpConstructor(root.Left,list)
     *list=append(*list,root.Val)
    helpConstructor(root.Right,list)
}


func (this *BSTIterator) Next() int {//下一个
    if this.HasNext(){
        val:=this.Val[0]
        this.Val=this.Val[1:]
        return val
    }
    return -1

}


func (this *BSTIterator) HasNext() bool {//是否有下一个
    if len(this.Val)==0{
        return false
    }
    return true

}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
