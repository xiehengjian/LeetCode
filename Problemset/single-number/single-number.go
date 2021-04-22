
// @Title: 只出现一次的数字 (Single Number)
// @Author: 846188037@qq.com
// @Date: 2021-03-22 17:55:02
// @Runtime: 16 ms
// @Memory: 6.1 MB

func singleNumber(nums []int) int {
    res:=0
    for i:=0;i<len(nums);i++{
        res=res^nums[i]
    }
    return res
}
