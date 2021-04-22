
// @Title: 寻找峰值 (Find Peak Element)
// @Author: 846188037@qq.com
// @Date: 2021-04-22 14:35:00
// @Runtime: 4 ms
// @Memory: 2.7 MB

func findPeakElement(nums []int) int {
    if len(nums)==1{
        return 0
    }
    if nums[0]>nums[1]{
        return 0
    }
    if nums[len(nums)-1]>nums[len(nums)-2]{
        return len(nums)-1
    }
    left,right:=1,len(nums)-2

    for left<=right{
        mid:=(left+right)/2
        if nums[mid]>nums[mid-1] && nums[mid]>nums[mid+1]{
            return mid
        }
        if nums[mid-1]<nums[mid] && nums[mid]<nums[mid+1]{
            left=mid+1
        }else{
            right=mid-1
        }
    }
    return -1

}
