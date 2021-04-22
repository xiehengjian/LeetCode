
// @Title: 跳跃游戏 (Jump Game)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 20:19:50
// @Runtime: 508 ms
// @Memory: 4.2 MB

func canJump(nums []int) bool {
	if len(nums)==0{
		return true
	}
	f:=make([]bool,len(nums))
	f[0]=true
	for i:=1;i<len(nums);i++{
		for j:=0;j<i;j++{
			if f[j]==true && nums[j]+j>=i{
				f[i]=true
				//break
			}
		}
	}
	return f[len(nums)-1]
}
