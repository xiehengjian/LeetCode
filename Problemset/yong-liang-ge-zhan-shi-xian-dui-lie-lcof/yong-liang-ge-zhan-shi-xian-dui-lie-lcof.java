
// @Title: 用两个栈实现队列 (用两个栈实现队列 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-03-31 17:59:08
// @Runtime: 58 ms
// @Memory: 46.4 MB

class CQueue {
    public Stack<Integer> stackPush;
    public Stack<Integer> stackPop;

    public CQueue() {
        stackPush=new Stack<Integer>();
        stackPop=new Stack<Integer>();

    }
    
    public void appendTail(int value) {
        stackPush.push(value);

    }
    
    public int deleteHead() {
        if(stackPop.empty() && stackPush.empty()){
            return -1;
        }else if (stackPop.empty()){
            while(!stackPush.empty()){
                stackPop.push(stackPush.pop());
            }
        }
        return stackPop.pop();

    }
}

/**
 * Your CQueue object will be instantiated and called as such:
 * CQueue obj = new CQueue();
 * obj.appendTail(value);
 * int param_2 = obj.deleteHead();
 */
