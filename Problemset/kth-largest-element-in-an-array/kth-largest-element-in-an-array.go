
// @Title: 数组中的第K个最大元素 (Kth Largest Element in an Array)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 09:23:04
// @Runtime: 8 ms
// @Memory: 3.4 MB

func findKthLargest(nums []int, k int) int {
    lo,hi:=0,len(nums)-1//首尾
    k=len(nums)-k //升序中的位置
    for lo<=hi{
        //一次划分，得到一个元素的确定位置
        p:=partition(nums,lo,hi)
        if p<k{
            //说明所求元素在p后面
            lo=p+1
        }else if p>k{
            //说明所求元素在p前面
            hi=p-1
        }else{
            //说明所求元素就是p
            return nums[p]
        }
    }
    return -1

}


func partition(nums []int, lo,hi int)int{
    if lo==hi{//待排序只有一个元素
        return lo
    }
    pivot:=nums[lo];//取左侧为轴
    i,j:=lo+1,hi//设定左右下标
    for{
        // 保证 nums[lo..i] 都小于 pivot
        for nums[i]<=pivot{
            i++
            if i>=hi{//走到头了
                break
            }
        }
        // 保证 nums[j..hi] 都大于 pivot
        for nums[j]>=pivot{
            j--
            if j<=lo{//走到头了
                break
            }
        }
        if i>=j{//两指针相遇结束本次划分
            break
        }
        // 如果走到这里，一定有：
        // nums[i] > pivot && nums[j] < pivot
        // 所以需要交换 nums[i] 和 nums[j]，
        // 保证 nums[lo..i] < pivot < nums[j..hi]
        nums[i],nums[j]=nums[j],nums[i]
    }
    nums[j],nums[lo]=nums[lo],nums[j]
    return j
}


