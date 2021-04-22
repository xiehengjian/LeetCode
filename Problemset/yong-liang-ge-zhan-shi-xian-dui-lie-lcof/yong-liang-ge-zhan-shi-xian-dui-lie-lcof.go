
// @Title: 用两个栈实现队列 (用两个栈实现队列 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 08:58:35
// @Runtime: 240 ms
// @Memory: 8.4 MB

type CQueue struct {
    Head []int
    Tail []int
}


func Constructor() CQueue {
    return CQueue{
        Head:[]int{},
        Tail:[]int{},
    }

}


func (this *CQueue) AppendTail(value int)  {
    this.Tail=append(this.Tail,value)

}


func (this *CQueue) DeleteHead() int {
    if len(this.Head)!=0{
        val:=this.Head[len(this.Head)-1]
        this.Head=this.Head[:len(this.Head)-1]
        return val
    }
    for len(this.Tail)!=0{
        this.Head=append(this.Head,this.Tail[len(this.Tail)-1])
        this.Tail=this.Tail[:len(this.Tail)-1]
    }
    if len(this.Head)!=0{
        val:=this.Head[len(this.Head)-1]
        this.Head=this.Head[:len(this.Head)-1]
        return val
    }
    return -1

}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
