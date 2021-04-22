
// @Title: 包含min函数的栈 (包含min函数的栈 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 13:54:13
// @Runtime: 20 ms
// @Memory: 7.9 MB

type MinStack struct {
        stack []int
    min []int

}


/** initialize your data structure here. */
func Constructor() MinStack {
        return MinStack{
        stack:make([]int,0),
        min:make([]int,0),
    }

}


func (this *MinStack) Push(x int)  {
        min:=this.Min()
    if x<=min{
        this.min=append(this.min,x)
    }
    this.stack=append(this.stack,x)

}


func (this *MinStack) Pop()  {
        if len(this.stack)==0{
        return 
    }
    if this.Min() == this.Top(){
        this.min=this.min[:len(this.min)-1]
    }
    this.stack=this.stack[:len(this.stack)-1]

}


func (this *MinStack) Top() int {
        if len(this.stack)==0{
        return 0
    }
    return this.stack[len(this.stack)-1]

}


func (this *MinStack) Min() int {
        if len(this.min)==0{
        return 1<<31
    }
    return this.min[len(this.min)-1]

}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
