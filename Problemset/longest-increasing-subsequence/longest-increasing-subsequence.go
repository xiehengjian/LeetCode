
// @Title: 最长递增子序列 (Longest Increasing Subsequence)
// @Author: 846188037@qq.com
// @Date: 2021-04-19 15:43:40
// @Runtime: 68 ms
// @Memory: 4 MB

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	//f[i]表示前i项里包含第i项的最长
	//  1  99  100  2  3 4
	n:=len(nums)
	if n==0{
		return 0
	}
	f:=make([]int,n+1)
	f[0]=0
	f[1]=1
	for i:=1;i<n+1;i++{
		f[i]=1
		for j:=0;j<i;j++{
            //所以对第i项来说，只要它比前面的第j项大，
            //那么至少f[i]>=f[j]+1
			if nums[j]<nums[i-1]{
				f[i]=max(f[i],f[j+1]+1)
			}
		}
	}
    fmt.Println(f)
	ans:=0
	for _,v:=range f{
		ans=max(ans,v)
	}
	return ans


}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
// @lc code=end
