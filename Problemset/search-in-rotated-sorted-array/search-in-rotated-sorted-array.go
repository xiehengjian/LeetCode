
// @Title: 搜索旋转排序数组 (Search in Rotated Sorted Array)
// @Author: 846188037@qq.com
// @Date: 2021-04-22 11:00:02
// @Runtime: 4 ms
// @Memory: 2.5 MB

func search(nums []int, target int) int {
    if len(nums)==0{
        return -1
    }
    left,right:=0,len(nums)-1
    for left<=right{
        mid:=left+(right-left)/2
        if nums[mid]==target{
            return mid
        }
        if nums[left]<=nums[mid]{
            if nums[left]<=target && target<=nums[mid]{
                right=mid-1
            }else{
                left=mid+1
            }

        }else{
            if nums[mid]<=target && target<=nums[right]{
                left=mid+1
            }else{
                right=mid-1
            }

        }
    }
    return -1

}
