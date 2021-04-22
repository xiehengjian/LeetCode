
// @Title: 跳跃游戏 II (Jump Game II)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 20:34:12
// @Runtime: 4 ms
// @Memory: 2.1 MB

func jump(nums []int) int {
	if len(nums)==0{
		return 0
	}
	n:=len(nums)
	f:=make([]int,n)
	f[0]=0
	for i:=1;i<n;i++{
		f[i]=math.MaxInt32
		for j:=0;j<i;j++{
			if f[j]<math.MaxInt32 && nums[j]+j>=i{
				fmt.Println(",")
				//fmt.Println()
				f[i]=min(f[j]+1,f[i])
			}
		}
	}
	fmt.Println(f)
	return f[n-1]
	

}

func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}
