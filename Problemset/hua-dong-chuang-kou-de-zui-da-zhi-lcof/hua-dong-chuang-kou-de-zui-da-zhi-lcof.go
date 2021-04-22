
// @Title: 滑动窗口的最大值 (滑动窗口的最大值 LCOF)
// @Author: 846188037@qq.com
// @Date: 2020-12-30 17:58:43
// @Runtime: 20 ms
// @Memory: 6.5 MB

func maxSlidingWindow(nums []int, k int) []int {
    qmax:=make([]int,0)
    res:=make([]int,0)
    for i:=0;i<len(nums);i++{
        for len(qmax)!=0 && nums[i]>nums[qmax[len(qmax)-1]]{
            qmax=qmax[:len(qmax)-1]
        }
        qmax=append(qmax,i)
        for qmax[0]<=i-k{
            qmax=qmax[1:]
        }
        if i>=k-1{
        res=append(res,nums[qmax[0]])
        }

    }
    return res
    
}
