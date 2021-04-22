
// @Title: 缺失的第一个正数 (First Missing Positive)
// @Author: 846188037@qq.com
// @Date: 2021-04-20 21:01:12
// @Runtime: 0 ms
// @Memory: 2.4 MB

func firstMissingPositive(nums []int) int {
    Map:=make(map[int]int)
    for i:=0;i<len(nums);i++{
        if nums[i]>0{
            Map[nums[i]]++
        }
    }
    for i:=1;i<len(Map)+1;i++{
        if _,ok:=Map[i];!ok{
            return i
        }
    }
    return len(Map)+1
}
