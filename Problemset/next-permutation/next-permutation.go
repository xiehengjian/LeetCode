
// @Title: 下一个排列 (Next Permutation)
// @Author: 846188037@qq.com
// @Date: 2021-04-22 08:59:19
// @Runtime: 4 ms
// @Memory: 2.5 MB

func nextPermutation(nums []int)  {
    if len(nums)==0 || len(nums)==1{
        return 
    }
    k,j:=0,0
    for i:=len(nums)-1;i>-1;i--{
        if i==0{
            reverse(nums)
            return
        }
        if nums[i]>nums[i-1]{
            k,j=i-1,i
            break
        }
    }

    m:=j
    for i:=j;i<len(nums);i++{
        if i==len(nums)-1 && nums[len(nums)-1]>nums[k]{
            m=i
            break
        }
        if nums[i]>nums[k] && nums[i+1]<=nums[k]{
            m=i
            break
        }
    }
    nums[k],nums[m]=nums[m],nums[k]
    
    reverse(nums[k+1:])
    
    

}

func reverse(nums []int)[]int{
   fmt.Println(nums)
    for i:=0;i<(len(nums))/2;i++{
        fmt.Println(len(nums)-i-1)
        nums[i],nums[len(nums)-i-1]=nums[len(nums)-i-1],nums[i]
    }
    fmt.Println(nums)
    return nums
}
