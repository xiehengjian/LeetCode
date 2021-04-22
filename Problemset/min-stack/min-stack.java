
// @Title: 最小栈 (Min Stack)
// @Author: 846188037@qq.com
// @Date: 2021-03-31 17:48:21
// @Runtime: 7 ms
// @Memory: 40.3 MB

class MinStack {

    /** initialize your data structure here. */
    private Stack<Integer> stackData;
    private Stack<Integer> stackMin;
    public MinStack() {
        this.stackData=new Stack<Integer>();
        this.stackMin=new Stack<Integer>();

    }
    
    public void push(int val) {
        if (this.stackMin.isEmpty()){
            this.stackMin.push(val);
        }else if(val<=this.getMin()){
            this.stackMin.push(val);
        }
        this.stackData.push(val);

    }
    
    public void pop() {
        if (this.stackData.isEmpty()){
            throw new RuntimeException("Your stack is a Empty");
        }
        int value=stackData.pop();
        if (value==this.getMin()){
            this.stackMin.pop();
        }
        //return value;

    }
    
    public int top() {
        if (this.stackData.isEmpty()){
            throw new RuntimeException("Your stack is empty");
        }
        return this.stackData.peek();

    }
    
    public int getMin() {
        if (this.stackMin.isEmpty()){
            throw new RuntimeException("Your Stack is empty");
        }
        return this.stackMin.peek();

    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(val);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */
