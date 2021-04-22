
// @Title: 最小栈 (Min Stack)
// @Author: 846188037@qq.com
// @Date: 2020-10-20 11:20:03
// @Runtime: 16 ms
// @Memory: 10 MB

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
    min:=this.GetMin()
    if x<=min{
        this.min=append(this.min,x)
    }
    this.stack=append(this.stack,x)

}


func (this *MinStack) Pop()  {
    if len(this.stack)==0{
        return 
    }
    if this.GetMin() == this.Top(){
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


func (this *MinStack) GetMin() int {
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
 * param_4 := obj.GetMin();
 */
