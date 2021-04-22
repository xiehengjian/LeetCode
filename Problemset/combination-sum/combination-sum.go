
// @Title: 组合总和 (Combination Sum)
// @Author: 846188037@qq.com
// @Date: 2021-04-20 21:22:52
// @Runtime: 4 ms
// @Memory: 3.1 MB

func combinationSum(candidates []int, target int) [][]int {
    res:=make([][]int,0)
    temp:=make([]int,0)
    //visited:=make([]int,len(candidates))
    backtrack(&res,temp,candidates,target,0)
    return res
    
}

func backtrack(res *[][]int,temp,candidates []int,target,start int){
    if sum(temp)>target{
        return
    }
    if sum(temp)==target{
        ans:=make([]int,len(temp))
        copy(ans,temp)
        *res=append(*res,ans)
        return
    }
    for i:=start;i<len(candidates);i++{
       
        temp=append(temp,candidates[i])
       
        backtrack(res,temp,candidates,target,i)
        temp=temp[:len(temp)-1]
       
    }
}

func sum(a []int)int{
    sum:=0
    for _,v:=range a{
        sum+=v
    }
    return sum
}
