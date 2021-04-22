
// @Title: 不同路径 (Unique Paths)
// @Author: 846188037@qq.com
// @Date: 2021-04-15 19:58:36
// @Runtime: 0 ms
// @Memory: 2 MB

func uniquePaths(m int, n int) int {
	if m==0 || n==0{
		return 0
	}
	f:=make([][]int,m)
	for i:=0;i<m;i++{
		f[i]=make([]int,n)
		f[i][0]=1
	}
	for i:=0;i<n;i++{
		f[0][i]=1
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			f[i][j]=f[i-1][j]+f[i][j-1]
		}
	}
	return f[m-1][n-1]

}
