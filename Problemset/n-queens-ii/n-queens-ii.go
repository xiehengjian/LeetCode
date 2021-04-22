
// @Title: N皇后 II (N-Queens II)
// @Author: 846188037@qq.com
// @Date: 2021-01-28 23:05:51
// @Runtime: 0 ms
// @Memory: 2 MB

func totalNQueens(n int) int {
    res:=0
    temp:=[]int{}
    backtrack(n,temp,&res)
    return res
}

func backtrack(n int,temp []int,res *int){
    if len(temp)==n{
        (*res)++
        return 
    }
    for i:=0;i<n;i++{
        if !isValid(i,temp){
            continue
        }
        temp=append(temp,i)
        backtrack(n,temp,res)
        temp=temp[:len(temp)-1]
        
    }

}

func isValid(i int,temp []int)bool{
    for k,v :=range temp{
            if i==v || i==v+len(temp)-k || i==v+k-len(temp){
                return false
            }
            //比如len(temp)==2,那就是要放第三个
            //那对于k=0,不能差两个,
            //对k=1就是不能差1个
            //不能差len(temp)-k个

        }
    return true
}
