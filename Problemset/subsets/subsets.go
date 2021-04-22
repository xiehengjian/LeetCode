
// @Title: 子集 (Subsets)
// @Author: 846188037@qq.com
// @Date: 2020-10-21 08:29:35
// @Runtime: 0 ms
// @Memory: 2.3 MB

func subsets(nums []int) [][]int {
    //保存最终结果
    result:=make([][]int,0)
    //保存中间结果
    list:=make([]int,0)
    backtrack(nums,0,list,&result)
    return result

}


func backtrack(nums []int,pos int,list []int,result *[][]int){
    ans:=make([]int,len(list))
    copy(ans,list)
    *result=append(*result,ans)
    for i:=pos;i<len(nums);i++{
        list=append(list,nums[i])
        backtrack(nums,i+1,list,result)
        list=list[0:len(list)-1]
    }
}
