//python中的暴力遍历搬过来的，但是Go没有in判断，所以又循环了一遍，速度比较慢
func findRepeatNumber(nums []int) int {
    for i:=0;i<len(nums);i++{
        for j:=0;j<len(nums[:i]);j++{
            if nums[i]==nums[:i][j]{
                return nums[i]
            }
        }
    }
    return -1
}
