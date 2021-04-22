
// @Title: 三角形最小路径和 (Triangle)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 16:41:02
// @Runtime: 4 ms
// @Memory: 3.3 MB

func minimumTotal(triangle [][]int) int {
	n:=len(triangle)
	f:=make([][]int,n)
	for i:=0;i<n;i++{
		f[i]=make([]int,i+1)
	}
	f[0][0]=triangle[0][0]
	for i:=1;i<n;i++{
		f[i][0]=f[i-1][0]+triangle[i][0]
		for j:=1;j<i;j++{
			f[i][j]=min(f[i-1][j],f[i-1][j-1])+triangle[i][j]
		}
		f[i][i]=f[i-1][i-1]+triangle[i][i]
	}
	ans:=math.MaxInt32
	for _,v:=range f[n-1]{
		ans=min(ans,v)
	}
	return ans
	

}

func min(x,y int)int{
	if x<y{
		return x
	}
	return y
}
