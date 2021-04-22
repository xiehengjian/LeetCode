
// @Title: 全排列 (Permutations)
// @Author: 846188037@qq.com
// @Date: 2021-01-17 13:28:44
// @Runtime: 0 ms
// @Memory: 2.7 MB

func permute(nums []int) [][]int {
    result:=make([][]int,0)
    list:=make([]int,0)
    visited:=make([]bool,len(nums))
    backtrack(nums,visited,list,&result)
    return result
}

func backtrack(nums []int,visited []bool,list []int, result *[][]int){
    if len(list)==len(nums){
        ans:=make([]int,len(list))
        copy(ans,list)
        *result=append(*result,ans)
        return 
    }

    for i:=0;i<len(nums);i++{
        if visited[i]{
            continue
        }
        list=append(list,nums[i])
        visited[i]=true
        backtrack(nums,visited,list,result)
        visited[i]=false
        list=list[0:len(list)-1]
    }
}
