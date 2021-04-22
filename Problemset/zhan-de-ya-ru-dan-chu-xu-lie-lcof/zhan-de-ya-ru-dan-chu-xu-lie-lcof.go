
// @Title: 栈的压入、弹出序列 (栈的压入、弹出序列 LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 14:10:38
// @Runtime: 8 ms
// @Memory: 3.8 MB

func validateStackSequences(pushed []int, popped []int) bool {
    stack := []int{}
    res := false
    // 模拟进栈
    for i:=0;i<len(pushed);i++{
        // 进栈
        stack = append(stack,pushed[i])
        // 出栈
        for popped[0]==stack[len(stack)-1]{
            popped = popped[1:]
            stack = stack[:len(stack)-1]
            // 防止slice越界
            if len(stack) == 0{
                break
            }
        }
    }
    if len(popped) == 0 {
        res = true
    }
    return res
}


