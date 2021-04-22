
// @Title: 二分查找 (Binary Search)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 11:36:09
// @Runtime: 40 ms
// @Memory: 6.7 MB

func search(nums []int, target int) int {
    start:=0
    end:=len(nums)-1
    for start<=end{
        mid:=start+(end-start)/2
        if nums[mid]==target{
            //end=mid
            return mid
        }else if nums[mid]<target{
            start=mid+1
        }else if nums[mid]>target{
            end=mid-1
        }
    }
    return -1


}
