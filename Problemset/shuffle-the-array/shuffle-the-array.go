
// @Title: 重新排列数组 (Shuffle the Array)
// @Author: 846188037@qq.com
// @Date: 2020-07-15 20:07:55
// @Runtime: 4 ms
// @Memory: 3.8 MB

func shuffle(nums []int, n int) []int {
    var newNums []int
    for i:=0;i<n;i++{
        newNums=append(newNums,nums[i],nums[n+i])
    }
    return newNums

}
