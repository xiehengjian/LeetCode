
// @Title: 三数之和 (3Sum)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 10:22:22
// @Runtime: 40 ms
// @Memory: 7.5 MB

func threeSum(nums []int) [][]int {
    n:=len(nums)
    sort.Ints(nums)
    res:=make([][]int,0)
    for first:=0;first<n;first++{
        if first>0 && nums[first]==nums[first-1]{
            continue
        }
        target:=-nums[first]
        second,third:=first+1,n-1
        for second<third{
            left,right:=nums[second],nums[third]
            sum:=left+right
            if sum<target{
                for second<third&& left==nums[second]{
                    second++
                }
            }else if sum>target{
                for second<third&& right==nums[third]{
                    third--
                }
            }else{
                res=append(res,[]int{nums[first],nums[second],nums[third]})
                 for second<third&& left==nums[second]{
                    second++
                }
                for second<third&& right==nums[third]{
                    third--
                }
            }
        } 
    }
     return res
}
