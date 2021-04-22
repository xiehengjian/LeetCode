
// @Title: 二叉树的序列化与反序列化 (Serialize and Deserialize Binary Tree)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 10:53:29
// @Runtime: 268 ms
// @Memory: 13.7 MB

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    l []string
    
}

func Constructor() Codec {
    return Codec{}
    
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var s string 
    res:=helpSerialize(root,s)
    fmt.Println(res)
    return res
}

func helpSerialize(root *TreeNode,s string)string{
    if root==nil{
        s+="NULL,"
        return s
    }
    s+=strconv.Itoa(root.Val)+","
    s=helpSerialize(root.Left,s)
    s=helpSerialize(root.Right,s)
    return s
}


// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode { 
    l:=strings.Split(data,",")
    for i:=0;i<len(l);i++{
        if l[i]!=""{
            this.l=append(this.l,l[i])
        }
    }
    fmt.Println(this.l)
    return this.helpDeserialize()
}

func (this *Codec)helpDeserialize()*TreeNode{
    if this.l[0]=="NULL"{
        this.l=this.l[1:]
        return nil 
    }
    val,_:=strconv.Atoi(this.l[0])
    root:=&TreeNode{Val:val}
    this.l=this.l[1:]
    root.Left=this.helpDeserialize()
    root.Right=this.helpDeserialize()
    return root

}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
