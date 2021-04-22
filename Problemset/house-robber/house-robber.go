
// @Title: 打家劫舍 (House Robber)
// @Author: 846188037@qq.com
// @Date: 2021-04-22 14:17:11
// @Runtime: 0 ms
// @Memory: 2.1 MB

func rob(nums []int) int {
    //f[i]第i个房间的最高金额
    f:=make([][]int,len(nums))
    for i:=0;i<len(nums);i++{
        f[i]=make([]int,2)
    }
    f[0][0]=0
    f[0][1]=nums[0]
    n:=len(nums)
    for i:=1;i<len(nums);i++{
        f[i][0]=max(f[i-1][0],f[i-1][1])
        f[i][1]=f[i-1][0]+nums[i]
    }

    return max(f[n-1][0],f[n-1][1])
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
