
// @Title: 0～n-1中缺失的数字 (缺失的数字  LCOF)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 16:29:14
// @Runtime: 20 ms
// @Memory: 6.1 MB

func missingNumber(nums []int) int {
    n:=0
    for i:=0;i<len(nums);i++{
        n=i
        if i!=nums[i]{  
            break
        }
        
    }
    if n==nums[len(nums)-1]{
        return n+1
    }
    return n

}
