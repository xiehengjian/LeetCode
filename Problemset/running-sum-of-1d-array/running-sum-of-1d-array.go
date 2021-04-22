
// @Title: 一维数组的动态和 (Running Sum of 1d Array)
// @Author: 846188037@qq.com
// @Date: 2020-07-14 20:53:18
// @Runtime: 4 ms
// @Memory: 2.7 MB

func runningSum(nums []int) []int {
    var runningSum []int
    var sum int
    for i:=0;i<len(nums);i++{
        for j:=0;j<i+1;j++{
            sum+=nums[j]   
        }
        runningSum=append(runningSum,sum) 
        sum=0
    }
    return runningSum
}
